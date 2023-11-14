package traffic

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/Layr-Labs/eigenda/common"
	"github.com/Layr-Labs/eigenda/common/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TrafficGenerator struct {
	Logger          common.Logger
	DisperserClient DisperserClient
	Config          *Config
}

func NewTrafficGenerator(config *Config) (*TrafficGenerator, error) {
	logger, err := logging.GetLogger(logging.DefaultCLIConfig())
	if err != nil {
		return nil, err
	}

	return &TrafficGenerator{
		Logger:          logger,
		DisperserClient: NewDisperserClient(config),
		Config:          config,
	}, nil
}

func (g *TrafficGenerator) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i := 0; i < int(g.Config.NumInstances); i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_ = g.StartTraffic(ctx)
		}()
		time.Sleep(g.Config.InstanceLaunchInterval)
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)
	<-signals

	cancel()
	wg.Wait()
	return nil
}

func (g *TrafficGenerator) StartTraffic(ctx context.Context) error {
	data := make([]byte, g.Config.DataSize)
	_, err := rand.Read(data)
	if err != nil {
		return err
	}

	ticker := time.NewTicker(g.Config.RequestInterval)
	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			if g.Config.RandomizeBlobs {
				_, err := rand.Read(data)
				if err != nil {
					return err
				}
			}
			err := g.sendRequest(ctx, data, 0)
			if err != nil {
				g.Logger.Error("failed to send blob request", "err:", err)
			}
		}
	}
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func (g *TrafficGenerator) sendRequest(ctx context.Context, data []byte, quorumID uint8) error {
	header := metadata.New(map[string]string{"x-forwarded-for": randSeq(32)})
	fmt.Println("sending request with header: ", header)
	grpc.SendHeader(ctx, header)
	outgoingCtx := metadata.NewOutgoingContext(ctx, header)
	ctxTimeout, cancel := context.WithTimeout(outgoingCtx, g.Config.Timeout)
	defer cancel()
	blobStatus, key, err := g.DisperserClient.DisperseBlob(ctxTimeout, data, quorumID, g.Config.QuorumThreshold, g.Config.AdversarialThreshold)
	if err != nil {
		return err
	}

	g.Logger.Info("successfully dispersed new blob,", "key", hex.EncodeToString(key), "status", blobStatus.String())
	return nil
}

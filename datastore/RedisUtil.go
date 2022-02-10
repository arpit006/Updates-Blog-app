package datastore

import (
	"github.com/go-redis/redis/v8"
	"log"
)

// RedisPipe created to execute Redis commands in a single pipeline (Bulk)
type RedisPipe struct {
	Pipe redis.Pipeliner
}

// StartPipeline starts the Redis Pipeline
func StartPipeline() *RedisPipe {
	client := GetRedisClientFactory()
	return &RedisPipe{Pipe: client.Pipeline()}
}

// ExecInPipeStrStr method over RedisPipe object! Executes in Pipeline
func (rp *RedisPipe) ExecInPipeStrStr(hash, key, value string) {
	rp.Pipe.HSet(ctx, hash, key, value)
}

// ExecInPipeStrInt method over RedisPipe object! Executes in pipeline
func (rp *RedisPipe) ExecInPipeStrInt(hash, key string, value int64) {
	rp.Pipe.HSet(ctx, hash, key, value)
}

// ExecInPipeIntInt method over RedisPipe object! Executes in pipeline
func (rp *RedisPipe) ExecInPipeIntInt(hash string, key, value int64) {
	rp.Pipe.HSet(ctx, hash, key, value)
}

// ExecInPipeIntStr method over RedisPipe object! Executes in pipeline
func (rp *RedisPipe) ExecInPipeIntStr(hash string, key int64, value string) {
	rp.Pipe.HSet(ctx, hash, key, value)
}

// ExecInPipeStrBytesArr method over RedisPipe object! Executes in pipeline
func (rp *RedisPipe) ExecInPipeStrBytesArr(hash, key string, value []byte) {
	rp.Pipe.HSet(ctx, hash, key, value)
}

func (rp *RedisPipe) LPushInPipeStrInt(key string, value int64) {
	rp.Pipe.LPush(ctx, key, value)
}

// ExecPipeline closes the pipeline. Performs bulk operations
func (rp *RedisPipe) ExecPipeline() error {
	_, err := rp.Pipe.Exec(ctx)
	if err != nil {
		log.Printf("Error in executing Redis Pipeline. Error is %s", err)
		return err
	}
	return nil
}

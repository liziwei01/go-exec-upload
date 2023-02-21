/*
 * @Author: liziwei01
 * @Date: 2022-03-21 22:35:37
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-24 23:08:03
 * @Description: file content
 */
package redis

import (
	"context"
	"testing"
)

func TestSet(t *testing.T) {
	ctx := context.Background()
	client, err := GetClient(ctx, "redis_idiary_cache")
	if err != nil {
		t.Error(err)
	}
	err = client.Set(ctx, "key", "value")
	if err != nil {
		t.Error(err)
	}
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	client, err := GetClient(ctx, "redis_idiary_cache")
	if err != nil {
		t.Error(err)
	}
	value, err := client.Get(ctx, "key")
	if err != nil {
		t.Error(err)
	}
	if value != "value" {
		t.Error(value)
	}
}

func TestDel(t *testing.T) {
	ctx := context.Background()
	client, err := GetClient(ctx, "redis_idiary_cache")
	if err != nil {
		t.Error(err)
	}
	err = client.Del(ctx, "key")
	if err != nil {
		t.Error(err)
	}
}

func TestExists(t *testing.T) {
	ctx := context.Background()
	client, err := GetClient(ctx, "redis_idiary_cache")
	if err != nil {
		t.Error(err)
	}
	exists, err := client.Exists(ctx, "key")
	if err != nil {
		t.Error(err)
	}
	if exists {
		t.Error(exists)
	}
}

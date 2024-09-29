// Copyright 2024 Authors of spidernet-io
// SPDX-License-Identifier: Apache-2.0

package draController

import (
	"context"
	"time"

	"github.com/spidernet-io/spiderpool/pkg/constant"
	"github.com/spidernet-io/spiderpool/pkg/election"
	clientset "github.com/spidernet-io/spiderpool/pkg/k8s/client/clientset/versioned"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/dynamic-resource-allocation/controller"
)

func StartController(ctx context.Context,
	leaderRetryElectGap time.Duration,
	spiderClientset clientset.Interface,
	kubeClient kubernetes.Interface,
	informerFactory informers.SharedInformerFactory,
	leader election.SpiderLeaseElector) error {

	driver := NewDriver(spiderClientset)
	controller := controller.New(ctx, constant.DRADriverName, driver, kubeClient, informerFactory)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case isLeader := <-leader.IsElected():
				if !isLeader {
					time.Sleep(leaderRetryElectGap)
					continue
				}
			}

			innerCtx, innerCancel := context.WithCancel(ctx)
			go func() {
				for {
					select {
					case <-innerCtx.Done():
						return
					case isLeader := <-leader.IsElected():
						if !isLeader {
							innerCancel()
							return
						}
					case <-ctx.Done():
						return
					}
					time.Sleep(leaderRetryElectGap)
				}
			}()

			informerFactory.Start(innerCtx.Done())
			controller.Run(1)
		}

	}()
	return nil
}

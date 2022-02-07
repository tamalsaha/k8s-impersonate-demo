package main

import (
	"context"
	"fmt"
	"k8s.io/client-go/rest"
	"reflect"

	//	"time"

	// 	"github.com/nats-io/nats.go"
	// 	"github.com/tamalsaha/nats-hop-demo/shared"
	// 	"github.com/tamalsaha/nats-hop-demo/transport"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	//	"k8s.io/klog/v2"
	"k8s.io/klog/v2/klogr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"
)

func NewClient() (client.Client, error) {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)

	ctrl.SetLogger(klogr.New())
	cfg := ctrl.GetConfigOrDie()
	cfg.QPS = 100
	cfg.Burst = 100

	mapper, err := apiutil.NewDynamicRESTMapper(cfg)
	if err != nil {
		return nil, err
	}

	cfg.Impersonate = rest.ImpersonationConfig{
		UserName: "system:serviceaccount:default:default",
		UID:      "",
		Groups:   nil,
		Extra:    nil,
	}

	return client.New(cfg, client.Options{
		Scheme: scheme,
		Mapper: mapper,
		//Opts: client.WarningHandlerOptions{
		//	SuppressWarnings:   false,
		//	AllowDuplicateLogs: false,
		//},
	})
}

func main__() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	c, err := NewClient()
	if err != nil {
		return err
	}

	var nodes core.NodeList
	err = c.List(context.TODO(), &nodes)
	if err != nil {
		panic(err)
	}
	for _, n := range nodes.Items {
		fmt.Println(n.Name)
	}
	return nil
}

type MyS struct {
}

func main() {
	fmt.Println(kind(MyS{}))
	fmt.Println(kind(&MyS{}))
}

func kind(v interface{}) string {
	return reflect.Indirect(reflect.ValueOf(v)).Type().Name()
}

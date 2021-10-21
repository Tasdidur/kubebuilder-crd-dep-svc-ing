/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"fmt"
	v1 "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	//"k8s.io/api/extensions/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"

	//"k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	//"k8s.io/apimachinery/pkg/api/errors"
	//"k8s.io/client-go/kubernetes"
	//"k8s.io/client-go/tools/cache"
	//"k8s.io/client-go/util/workqueue"
	//"time"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
	//appslister "k8s.io/client-go/listers/apps/v1"
	//corev1lister "k8s.io/client-go/listers/core/v1"
	kapiv1 "github.com/Tasdidur/kubebuilderProject/kbCrd/api/v1"
	netv1 "k8s.io/api/networking/v1"
	//lister "github.com/Tasdidur/kubebuilderProject/kbCrd"
	//kubeinformers "k8s.io/client-go/informers"
)

// KcrdReconciler reconciles a Kcrd object
type KcrdReconciler struct {
	client.Client
	Scheme *runtime.Scheme

	//kubeclientset   kubernetes.Interface
	//sampleclientset clientset.Interface
	//
	//deploymentsLister appslister.DeploymentLister
	//deploymentsSynced cache.InformerSynced
	//
	//serviceLister corev1lister.ServiceLister
	//serviceSynced cache.InformerSynced
	//
	//ingressLister netv1.IngressLister
	//ingressSynced cache.InformerSynced
}

//+kubebuilder:rbac:groups=kapi.tapi.com,resources=kcrds,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=kapi.tapi.com,resources=kcrds/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=kapi.tapi.com,resources=kcrds/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Kcrd object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.8.3/pkg/reconcile
func (r *KcrdReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	Log := log.FromContext(ctx)

	// your logic here

	kcrd := &kapiv1.Kcrd{}
	err := r.Client.Get(ctx, req.NamespacedName, kcrd)
	fmt.Println(kcrd.Name)

	getDep := &v1.Deployment{}
	err = r.Client.Get(ctx, req.NamespacedName, getDep)
	if errors.IsNotFound(err) {
		fmt.Println("new deployment creating... +++++++")
		if err = r.Client.Create(ctx, newDeployment(kcrd)); err != nil {
			Log.Error(err, "error creating  deployments ++++++")
		} else {
			fmt.Println(kcrd.Name + "-dep " + "created +++++++++")
		}
	}

	getSvc := &v12.Service{}
	err = r.Client.Get(ctx, req.NamespacedName, getSvc)
	if errors.IsNotFound(err) {
		fmt.Println("new service creating... +++++++")
		if err = r.Client.Create(ctx, newService(kcrd)); err != nil {
			Log.Error(err, "error creating  service ++++++")
		} else {
			fmt.Println(kcrd.Name + "-svc " + "created +++++++++")
		}
	}

	getIng := &netv1.Ingress{}
	err = r.Client.Get(ctx, req.NamespacedName, getIng)
	if errors.IsNotFound(err) {
		fmt.Println("new ingress creating... +++++++")
		if err = r.Client.Create(ctx, newIngress(kcrd)); err != nil {
			Log.Error(err, "error creating  ingress ++++++")
		} else {
			fmt.Println(kcrd.Name + "-ing " + "created +++++++++")
		}
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *KcrdReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&kapiv1.Kcrd{}).
		Owns(&v1.Deployment{}).
		Owns(&v12.Service{}).
		Owns(&netv1.Ingress{}).
		Complete(r)
}

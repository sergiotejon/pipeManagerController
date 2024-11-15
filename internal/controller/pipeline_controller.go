/*
Copyright 2024.

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

package controller

import (
	"context"
	"fmt"
	"k8s.io/apimachinery/pkg/api/errors"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	pipemanagerv1alpha1 "github.com/sergiotejon/pipeManagerController/api/v1alpha1"
)

// PipelineReconciler reconciles a Pipeline object
type PipelineReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=pipemanager.sergiotejon.github.io,resources=pipelines,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=pipemanager.sergiotejon.github.io,resources=pipelines/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=pipemanager.sergiotejon.github.io,resources=pipelines/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Pipeline object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.19.0/pkg/reconcile
func (r *PipelineReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	var pipeline pipemanagerv1alpha1.Pipeline
	if err := r.Get(ctx, req.NamespacedName, &pipeline); err != nil {
		if errors.IsNotFound(err) {
			log.Info("Pipeline resource not found. Ignoring since object must be deleted.")
			return ctrl.Result{}, nil
		}
		log.Error(err, "Failed to get PushMain.")
		return ctrl.Result{}, err
	}

	// TODO: Your logic here (Next an example)

	log.Info("Pipeline spec", "spec", pipeline.Spec)

	// TODO: Normalize
	fmt.Println("Normalize... Under construction.")
	// -- Read spec from k8s object (only one pipeline)
	// -- Refactor Normalize to work only with one pipeline
	// Normalize the pipelines
	//pipelines, err := normalize.Normalize(rawPipelines)
	//if err != nil {
	//	logging.Logger.Error("Error normalizing pipelines", "msg", err)
	//	os.Exit(ErrCodeNormalize)
	//}

	// Aquí puedes iterar sobre las tareas y gestionar cada una
	for taskName, task := range pipeline.Spec.Tasks {
		log.Info("Processing Task", "TaskName", taskName, "Description", task.Description)
		for _, step := range task.Steps {
			log.Info("  Step", "Name", step.Name, "Image", step.Image)
			// Implementa la lógica para manejar cada step
			// Por ejemplo, crear Pods o Jobs según los steps
		}
	}

	// Actualizar el estado si es necesario
	// pushMain.Status.SomeField = "SomeValue"
	// if err := r.Status().Update(ctx, &pushMain); err != nil {
	//     logger.Error(err, "Failed to update PushMain status")
	//     return ctrl.Result{}, err
	// }

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PipelineReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&pipemanagerv1alpha1.Pipeline{}).
		Named("pipeline").
		Complete(r)
}

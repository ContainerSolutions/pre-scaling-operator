package resources

import (
	"context"
	v1 "k8s.io/api/apps/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func DeploymentLister(ctx context.Context, _client client.Client, namespace string, OptInLabel map[string]string) (v1.DeploymentList, error) {

	deployments := v1.DeploymentList{}
	err := _client.List(ctx, &deployments, client.MatchingLabels(OptInLabel), client.InNamespace(namespace))
	if err != nil {
		return v1.DeploymentList{}, err
	}
	return deployments, nil

}

func DeploymentGetter(ctx context.Context, _client client.Client, req ctrl.Request) (v1.Deployment, error) {

	deployment := v1.Deployment{}
	err := _client.Get(ctx, req.NamespacedName, &deployment)
	if err != nil {
		return v1.Deployment{}, err
	}
	return deployment, nil

}

func DeploymentScaler(ctx context.Context, _client client.Client, deployment v1.Deployment, replicas int32) error {

	deployment.Spec.Replicas = &replicas
	err := _client.Update(ctx, &deployment, &client.UpdateOptions{})
	if err != nil {
		return err
	}

	return nil
}

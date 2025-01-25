// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	network "github.com/felixz92/provider-hcloud/internal/controller/hcloud/network"
	server "github.com/felixz92/provider-hcloud/internal/controller/hcloud/server"
	providerconfig "github.com/felixz92/provider-hcloud/internal/controller/providerconfig"
	subnet "github.com/felixz92/provider-hcloud/internal/controller/subnet/subnet"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		network.Setup,
		server.Setup,
		providerconfig.Setup,
		subnet.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}

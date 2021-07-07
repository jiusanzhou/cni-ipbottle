/*
 * Copyright (c) 2021 wellwell.work, LLC by Zoe
 *
 * Licensed under the Apache License 2.0 (the "License");
 * You may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"fmt"
	"log"

	"github.com/containernetworking/cni/pkg/skel"
	cniversion "github.com/containernetworking/cni/pkg/version"

	"go.zoe.im/cni-ipbottle/cmd"
	"go.zoe.im/cni-ipbottle/pkg/skelplugin"
	"go.zoe.im/x/cli"
	"go.zoe.im/x/version"
)

func main() {
	sp := skelplugin.New()

	cmd.Register()

	cmd.Option(
		// cli.GlobalConfig(svr.Config, cli.WithConfigName()),
		cli.Run(func(c *cli.Command, args ...string) {
			skel.PluginMain(
				sp.Add, sp.Check, sp.Del,
				cniversion.All,
				fmt.Sprintf("%s %s", c.Name(), version.Get().GitVersion),
			)
		}),
	)

	if err := cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}

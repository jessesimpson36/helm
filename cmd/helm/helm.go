/*
Copyright The Helm Authors.

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

package main // import "helm.sh/helm/v4/cmd/helm"

import (
	"fmt"
	chart "helm.sh/helm/v4/pkg/chart/v2"
	chartutil "helm.sh/helm/v4/pkg/chart/v2/util"
	"helm.sh/helm/v4/pkg/engine"
	"log"
	// Import to initialize client auth plugins.
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

const rawValues = `
title: Hello world
author: Jane Doe
date: Apr 8, 2025
intro: Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas.
listicle:
  - text: One thing
    link: https://one.example.com
  - text: Another thing
    link: https://two.example.com
  - text: Yet another
    link: https://three.example.com
body: |
  <p><strong>Pellentesque habitant morbi tristique</strong> senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. <em>Aenean ultricies mi vitae est.</em> Mauris placerat eleifend leo. Quisque sit amet est et sapien ullamcorper pharetra. Vestibulum erat wisi, condimentum sed, <code>commodo vitae</code>, ornare sit amet, wisi. Aenean fermentum, elit eget tincidunt condimentum, eros ipsum rutrum orci, sagittis tempus lacus enim ac dui. <a href="#">Donec non enim</a> in turpis pulvinar facilisis. Ut felis.</p>

  <h2>Header Level 2</h2>

  <ol>
      <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
      <li>Aliquam tincidunt mauris eu risus.</li>
  </ol>

  <blockquote><p>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vivamus magna. Cras in mi at felis aliquet congue. Ut a est eget ligula molestie gravida. Curabitur massa. Donec eleifend, libero at sagittis mollis, tellus est malesuada tellus, at luctus turpis elit sit amet quam. Vivamus pretium ornare est.</p></blockquote>

  <h3>Header Level 3</h3>

  <ul>
      <li>Lorem ipsum dolor sit amet, consectetuer adipiscing elit.</li>
      <li>Aliquam tincidunt mauris eu risus.</li>
  </ul>

  <pre><code>
  #header h1 a {
    display: block;
    width: 300px;
    height: 80px;
  }
  </code></pre>
footer: Copyright 2025 bad idea for Helm chart, Inc.

`

func main() {

	c := &chart.Chart{
		Metadata: &chart.Metadata{Name: "test", APIVersion: "v2", Version: "0.1.0"},
		Templates: []*chart.File{
			{Name: "templates/test.yaml", Data: []byte(
				`html: |
  <!DOCTYPE html>
  <html lang="en">
    <head>
      <meta charset="utf-8">
      <title>{{ .Values.title }}</title>
      <link rel="stylesheet" href="style.css">
      <script src="script.js"></script>
    </head>
    <body>
      <h1>{{ .Values.title }}</h1>
      <p><b>By {{ .Values.author }}</b> {{ .Values.date }}</p>
      {{ toYaml .Values.body | trimPrefix "|" | nindent 4 }}
      <p>{{ .Values.footer }}<p>
    </body>
  </html>`),
			},
		}}

	v, _ := chartutil.ReadValues([]byte(rawValues))

	vals := map[string]interface{}{
		"Values": v.AsMap(),
	}
	out, _ := new(engine.Engine).Render(c, vals)
	templateString := out["test/templates/test.yaml"]
	fmt.Println(templateString)
}

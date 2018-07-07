package main

var ddlTempl = `
metadata    :name        => "{{.Name}}",
            :description => "Choria Management Backplane",
            :author      => "R.I.Pienaar <rip@devco.net>",
            :license     => "Apache-2.0",
            :version     => "{{.Version}}",
            :url         => "https://choria.io/",
            :timeout     => 10

{{if .Health}}
action "health", :description => "Checks the health of the managed service" do
    output :result,
            :description => "The result from the check method",
            :display_as => "Result"

    output :healthy,
            :description => "Status indicator for the checked service",
            :display_as => "Healthy",
            :default => false

    summarize do
        aggregate summary(:healthy)
    end   
end
{{end}}

{{if .Stop}}
action "stop", :description => "Stops the managed service" do
    output :delay,
            :description => "How long after running the action the shutdown will be initiated",
            :display_as => "Delay"
end
{{end}}

{{if .Pause}}
["info", "pause", "resume", "flip"].each do |act|
    action act, :description => act do
        display :always

        output :paused,
               :description => "Circuit Breaker pause state",
               :display_as => "Paused"

        if act == "info"
            output :version,
                   :description => "Service Version",
                   :display_as => "Version"

            output :facts,
                   :description => "Instance Facts",
                   :display_as => "Facts"
        end

        summarize do
            aggregate summary(:version) if act == "info"
            aggregate summary(:paused)
        end
    end
end
{{end}}
`
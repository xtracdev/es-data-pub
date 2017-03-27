## ECS Task Execution Notes

To execute the task, first customize the particulars in 
`task-def-skeleton.json` to create `task-def.json`. Once this has
been done, register the task definition:

<pre>
aws ecs register-task-definition --cli-input-json file://$PWD/task-def.json
</pre>

Once the task has been defined in ECS, you can then run it:

<pre>
aws ecs run-task --cluster DemoCluster --task-definition pubevents
</pre>

You can status the task using the task arn returned by the run-task
command.

<pre>
aws ecs describe-tasks --cluster DemoCluster --tasks arn:aws:ecs:us-west-1:nnnnnnnn:task/69b799b1-5c3c-4464-a93e-765f23189bd9
</pre>
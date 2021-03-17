# Pre-Scaling Kubernetes Operator

Built out of necessity, the Operator helps pre-scale applications in anticipation of load.

At its core, it manages a cluster / namespace "state", 
which triggers scaling of applications to the desired replicas for that specific state.

For example, a cluster could have 2 states, `business-as-usual`, and `peak`.

Applications can configure the replicas needed for each state.
 
When a cluster is switched from one state to the other, 
the operator will take care of scaling all of the applications to the necessary replicas,
as set by developers through annotations.

Please check the `/docs` folder for documentation related to the architecture of the PreScale Operator 

# Multi Stage Deployments with Golang

## Kubernetes: A quick review

* Kubernetes runs your workload by placing containers into Pods to run on Nodes. A node may be a virtual or physical machine, depending on the cluster. Each node contains the services necessary to run Pods, managed by the control plane.

* A Kubernetes Pod is a group of one or more Containers, tied together for the purposes of administration and networking.

* Kubernetes Deployment checks on the health of your Pod and restarts the Pod's Container if it terminates. Deployments are the recommended way to manage the creation and scaling of Pods.

![kubeoverview](./images/kubeoverview.png)

Essentially, for our microservices framework to work, we need to have **NODES** that have **PODS** that control **CONTAINERS**. We'll need to use **minikube** to create and run our deployment (which is how we control our containers). We're going to do a basic deployment here of a basic **image**. 

Now- _normally_ we want images in **minikube** to come from a repository. Normally minikube would pull images from a docker repo where you had pushed the image to (that repo could be aws, git, dockerhub, a private repo...wherever). 

In our case, because we're running local versions of minikube and we don't want to go to the trouble of creating accounts and repositories for our docker images, we're going to instead use our local system. This requires a tiny bit of hackiness, admittedly.

FIRST OFF let's make our initial image!

![imagefire](./images/imagefire.jpeg)

Navigate into the **module_03/lab_03** directory.

**Side note:** In bad films there's a pretty common trope where one character will say to another "AS YOU KNOW....{exposition}" as a way to explain things to the audience. 
I always hated that as obviously people don't speak like that ("As you know...red lights mean stop").
BUT...I'm going to AS YOU KNOW you guys here and say: 
AS YOU KNOW- we `docker build` an image with the `docker build` command...so that's the first thing we're going to do now. Go to your command line and type in:

`docker build -t myproject/html-server-image:v1 .`.

![asyouknow](./images/asyouknow.png)

So now we have an image called **html-server-image** with the version **v1**. 
Let's make sure the image was built with:

`docker images`
Is your **html-server-image** there? (you can also run `docker images | grep html-server`)

Okay...now let's see if we can see the image **inside** minikube. Run the following on the command line:

`minikube ssh`

![images](./images/minikubessh.png)

Okay...let's run `docker images` from in here and see what we see!

Do you see your **html-server-image**? 
Assuming you don't (and you shouldn't) we'll need to somehow point our docker environment to minikube...which we will do now. Let's get out of ssh with a `exit` and go back to our basic command line. 

Once we're back in our local command line let's run this command:
`eval $(minikube docker-env)`

Now- that should switch our image to allow us to build "inside minikube". Run the following to check:

`docker images`

Do you see all of the `k8s.gcr` projects now in your "local"? Excellent! We have now shifted to our environment working inside of kubernetes and we're ready to rebuild our image! 
So let's do that...

Run `docker build -t myproject/html-server-image:v1 .` again.

And again, once everything is run, let's run `minikube ssh` and see what's in there with `docker images`.

Do you see your **myproject/html-server-image** sitting there in the repository list? 
AWESOME! Now we want to switch the discovery policy as Kubernetes usually checks to pull repositories by default. `exit` from your minikube and run 

`kubectl run practicehtml --image=myproject/html-server-image --image-pull-policy=Never --port=8761`

And from here let's take a look at our pods with `kubectl get pods`
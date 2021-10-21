# kubebuilder-crd-dep-svc-ing

### create cluster to run ingress controller and set the dns resolver
  <pre>
   $ kind create cluster --config clust.yaml 
   $ sudo nano /etc/hosts 
            add this -> 127.0.0.1   tasdid2.com
  </pre>

### add ingress controller
  <pre>
   $ kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml
  </pre>

### prepare your workdirectory with kubebuilder
  <pre>
  $ kubebuilder init --domain tapi.com
  </pre>
  the above command will create the following structure : 

  ![alt text](https://github.com/Tasdidur/kubebuilder-crd-dep-svc-ing/blob/c621904d162902026b47b91b8182ab722499a91b/images/Screenshot%20from%202021-10-21%2012-42-19.png?raw=true)

  then apply the following command :
  <pre>
  $ kubebuilder create api --group kapi --version v1 --kind Kcrd
  </pre>
  resulting into : 

  ![alt text](https://github.com/Tasdidur/kubebuilder-crd-dep-svc-ing/blob/c621904d162902026b47b91b8182ab722499a91b/images/Screenshot%20from%202021-10-21%2012-44-10.png)

# make the Crd
<pre>
  to make the crd we will follow the steps :
    1. edit the types.go
    2. generate the respective yaml
    3. apply the generated yaml to the cluster and introduce the crd to the world
    4. apply a test yaml
</pre>
### edit the types.go
  * in kubebuilder the types.go file exists as /api/v1/kcrd_types.go
  * get that and make edit 
### generate the respective yaml
  * in kubebuilder we dont need to run the command "$ controller-gen rbac:roleName=xt crd paths=./... output:crd:dir=sampleYaml output:stdout"\
    as we did in the sample controller style. kubebuilder provides a makefile.
  * run the following command :\
    $ make manifests
  * this will create the sample yaml named "kapi.tapi.com_kcrds.yaml" in the directory /config/crd/bases/
### introduce the crd to the world
  * run the command :\
    $ make install 
  * the command will apply the sample yaml in the cluster and introduce the crd 
### apply the test yaml
  * there is a yaml testkb.yaml
  * run the command :\
    $ kubectl apply -f testkb.yaml
  * if everything is ok, then you will see the msg that a new crd has been created

  
# make and run the controller
  * edit the kcrd_controller.go in the controller folder
  * edit the mgr option ; add SyncPeriod ; set it to a lesser time (default is 10h); so your conciler will repeat faster
  * run commands :\
    $ make build\
    $ make run
  * enjoy the show

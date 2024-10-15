#!/bin/bash

function print_logo() {
    echo -e "$blue
            

 ██▀███   ▒█████   ██▀███  
▓██ ▒ ██▒▒██▒  ██▒▓██ ▒ ██▒
▓██ ░▄█ ▒▒██░  ██▒▓██ ░▄█ ▒
▒██▀▀█▄  ▒██   ██░▒██▀▀█▄  
░██▓ ▒██▒░ ████▓▒░░██▓ ▒██▒
░ ▒▓ ░▒▓░░ ▒░▒░▒░ ░ ▒▓ ░▒▓░
  ░▒ ░ ▒░  ░ ▒ ▒░   ░▒ ░ ▒░
  ░░   ░ ░ ░ ░ ▒    ░░   ░ 
   ░         ░ ░     ░     
    _             _    __     _                    
   (_)   ____    (_)  / /_   (_) ____  ___    _____
  / /   / __ \  / /  / __/  / / /_  / / _ \  / ___/
 / /   / / / / / /  / /_   / /   / /_/  __/ / /    
/_/   /_/ /_/ /_/   \__/  /_/   /___/\___/ /_/     

    "
    echo -e "$clear"
}

function print_help() {
    # Display Help
    echo -e "$yellow"
    echo "Syntax: ./ror-cluster-initializer.sh [create|c|help|h]"
    echo
    echo "options:"
    echo "  create  alias: c    Create a local cluster with k3d and docker-desktop"
    echo "  help    alias: h    Print this Help"
    echo ""
    echo "dependencies: docker, k3d and kubectl"
    echo ""
    now=$(date)
    printf "Current date and time in Linux %s\n" "$now"
    echo ""
    echo -e "$clear"
}

clear

yellow='\033[0;33m'
clear='\033[0m'
blue='\033[0;34m'
red='\033[0;31m'

spinner()
{
    local pid=$!
    local delay=0.25
    local spinstr='|/-\'
    while [ "$(ps a | awk '{print $1}' | grep $pid)" ]; do
        local temp=${spinstr#?}
        printf "$blue [%c]  " "$spinstr"
        local spinstr=$temp${spinstr%"$temp"}
        sleep $delay
        printf "\b\b\b\b\b\b"
    done
    printf "    \b\b\b\b"
    echo -e "$clear"
}

function prerequisites() {
  if ! command -v $1 1> /dev/null
  then
      echo -e "$red 🚨 $1 could not be found. Install it! 🚨"
      exit
  fi
}

function get_cluster_parameter() {
    prerequisites docker
    prerequisites k3d
    prerequisites kubectl
    prerequisites helm

    echo -e "$clear"
    read -p "Enter the cluster name: " cluster_name
    read -p "Enter number of agents (1 or more): " agent_number
    #read -p "Enter the api port number (over 1024 but less than 65034, and not used from before): " random_port
    
    random_port=$(( ( RANDOM % 500 )  + 10000 ))
    echo port: $random_port
    extra=1
    loadbalancer_port=$(($extra + $random_port))

    # Uses netstat to check if a port is in ute, empty str if not in use
    port_used=$(netstat -tulpn 2>/dev/null | grep LISTEN | grep $random_port)
    loadbalancer_used=$(netstat -tulpn 2>/dev/null | grep LISTEN | grep $loadbalancer_port)

    while [ ! -z "${port_used}" ] || [ ! -z "${loadbalancer_used}" ]; do
        random_port=$(awk -v seed=$RANDOM 'begin{srand(seed);x=10000 + rand() * 5000; print int(x)}')
        echo port: $random_port
        extra=1
        loadbalancer_port=$(($extra + $random_port))
        port_used=$(netstat -tulpn 2>/dev/null | grep LISTEN | grep $random_port)
        loadbalancer_used=$(netstat -tulpn 2>/dev/null | grep LISTEN | grep $random_port)
    done

    read -p "Install ArgoCD and NHN-Tooling mock? (y/yes | n/no): " install_extra

    echo -e "$yellow
    k3d cluster create $cluster_name --api-port 127.0.0.1:$random_port -p \"$loadbalancer_port:80@loadbalancer\" --agents $agent_number"

    echo -e "$yellow
    Install ArgoCD and NHN-Tooling mock?: $install_extra"
    
    echo -e "$clear"
    read -p "Looks ok (y | yes)? " ok

    if [ "$ok" == "yes" ];then
            echo "Excellent"
            create_cluster        
        elif [ "$ok" == "y" ]; then
            echo "Good"
            create_cluster        
        else
            echo "That is bad ... quitting"
            exit 0
    fi
}

function install_extra(){
    echo -e "$yellow
    Create ArgoCD namespace
    "        
    (kubectl create namespace argocd|| 
    { 
        echo -e "$red 
        Could not namespace argocd in cluster ...
        "
        exit 1
    }) & spinner

    echo -e "$yellow
    Installing ArgoCD
    "
    (kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml|| 
    { 
        echo -e "$red 
        Could not install argocd into cluster  ...
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Setup NHN-Tooling mockup
    "
    (kubectl apply -n argocd -f ./argocd-proj-nhn-tooling.yaml|| 
    { 
        echo -e "$red 
        Could not add nhn-tooling ...
        "
        exit 1
    }) & spinner

    (kubectl apply -n argocd -f ./k8s-nhn-tooling.yaml|| 
    { 
        echo -e "$red 
        Could not install nhn-tooling to argocd in cluster ...
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Setup ror-agent mockup
    "

    (kubectl apply -f ./ror-namespace.yaml|| 
    { 
        echo -e "$red 
        Could not create ror namespace ...
        "
        exit 1
    }) & spinner

    (kubectl apply -n ror -f ./ror-configmap.yaml|| 
    { 
        echo -e "$red 
        Could not add configmap for ror ...
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Adding kyverno to helm repo
    "
    (helm repo add kyverno https://kyverno.github.io/kyverno/|| 
    { 
        echo -e "$red 
        Could not add kyverno to helm repo ...
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Creating kyverno namespace
    "
    (kubectl create namespace kyverno|| 
    { 
        echo -e "$red 
        Could not create kyverno namespace
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Install kyverno with helm
    "
    (helm upgrade --install kyverno --namespace kyverno kyverno/kyverno|| 
    { 
        echo -e "$red 
        Could not create kyverno with helm
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Install kyverno policies with helm
    "
    (helm upgrade --install kyverno-policies --namespace kyverno kyverno/kyverno-policies|| 
    { 
        echo -e "$red 
        Could not create kyverno policies with helm
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Install non root kyverno policies 
    "
    (kubectl apply -f policies/non-root.yaml|| 
    { 
        echo -e "$red 
        Could not create kyverno policies with helm
        "
        exit 1
    }) & spinner
}

function create_cluster() {
    if [ -z $cluster_name ] || [ $random_port -lt 1024 ] || [ $agent_number -lt 1 ] || [ $agent_number == $loadbalancer_port ]; then
        echo "Not all parameters good ... quitting"
        exit 1
    fi

    echo -e "$yellow
    Creating k3d cluster
    "

    echo -e "$clear"
    (k3d cluster create $cluster_name --api-port 127.0.0.1:$random_port -p "$loadbalancer_port:80@loadbalancer" --agents $agent_number || 
    { 
        echo -e "$red 
        Could not create cluster ...
        "
        exit 1
    }) & spinner

    (kubectl apply -f ./cluster-namespace.yaml|| 
    { 
        echo -e "$red 
        Could not create cluster namespace ...
        "
        exit 1
    }) & spinner

    echo -e "$yellow
     Installing ROR Tasks CRD (if exists)
    "
    CRDFILE=../../cmd/operator/config/crd/bases/ror.nhn.no_tasks.yaml
    if test -f "$CRDFILE"; then
        echo "ROR Task CRD exists."
        (kubectl apply -f $CRDFILE|| 
        { 
            echo -e "$red 
            Could not install ROR Operator (Task) CRD ...
            "
            exit 1
        }) & spinner
    else
        echo -e "$red
        Could not install ROR Operator (Task) CRD ...
        "
    fi

    echo -e "$yellow
     Setup ror-operator
    "
    (kubectl apply -f ./ror-operator.yaml|| 
    { 
        echo -e "$red 
        Could not create ror-operator clusterRoleBinding ...
        "
        exit 1
    }) & spinner

    if [ "$install_extra" == "yes" ];then
            install_extra
        elif [ "$install_extra" == "y" ]; then
            install_extra
    fi

    kubectl create namespace ror

    (k3d kubeconfig get -a > "k3d-$cluster_name.config" || 
    { 
        echo -e "$red
        Could not save
        "
        exit 1
    }) & spinner

    echo -e "$yellow
    Done creating k3d-cluster, config file saved to $red k3d-$cluster_name.config $yellow in current directory
    "
    echo -e "$yellow
    To delete cluster, type: $red k3d cluster delete $cluster_name $clear
    "
    echo -e "$clear"
}

while (($#)); do
   case $1 in
        create|c) # create cluster
            print_logo
            get_cluster_parameter
            exit
            ;;
        help|h) # display Help
            print_logo
            print_help
            exit;;
        *) # Invalid option
            echo -e "$red
            Error: Invalid option
            $clear
            "
            
            exit;;
   esac
done

print_logo
print_help

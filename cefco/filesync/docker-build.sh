#set -ex
if [ $# != 5 ] ; then 
 echo "please input args:
 	arg 1: image name (image name must be lowercase,eg: filereader)
	arg 2: image version（Tag）
	arg 3: pkg name (package must be in pkg directories )
	arg 4: system (centos:7.5 or alpine or debian)
	arg 5: cpu arch (amd64 or arm64)"
 exit 1; 
 fi
 
ROOT=$(cd $(dirname "$BASH_SOURCE") ; pwd)

NAME=$1
var=$3
PROCESS_NAME="${var%%-*}"
ARCH=$5

#alpine
#centos:7.5
#debian

BASE_OS=$4

if [[ $BASE_OS = "alpine" ]] ; then
	if [[ $ARCH = "amd64" ]] ; then
                BASE_OS_URL="registry-jinan-lab.inspurcloud.cn/library/os/inspur-alpine-3.10:5.0.0"
        elif [[ $ARCH = "arm64" ]] ;then
                BASE_OS_URL="registry-jinan-lab.inspurcloud.cn/library/os/inspur-alpine-3.10-arm64:5.0.0"
        else
                echo "Not support System : $BASE_OS"
        	exit
	fi
elif [[ $BASE_OS = "centos:7.5" ]] ;then 
	if [[ $ARCH = "amd64" ]] ; then
                BASE_OS_URL="registry-jinan-lab.inspurcloud.cn/library/os/inspur-centos-7:5.2.0"
        elif [[ $ARCH = "arm64" ]] ;then
                BASE_OS_URL="registry-jinan-lab.inspurcloud.cn/library/os/inspur-centos-7-arm64:5.0.0"
        else
                echo "Not support System : $BASE_OS"
        	exit
	fi
elif [[ $BASE_OS = "debian" ]] ;then
	if [[ $ARCH = "arm64" ]] ; then
    		BASE_OS_URL="registry-jinan-lab.inspurcloud.cn/library/os/inspur-debian-stretch-arm64:5.0.0"
	elif [[ $ARCH = "amd64" ]] ;then
    		BASE_OS_URL="registry-jinan-lab.inspurcloud.cn/library/os/inspur-debian-stretch:5.0.0"
	else
   		echo "Not support System : $BASE_OS"
   		exit
	fi
else
   echo "Not support System : $BASE_OS"
   exit

fi
BUILD_ARGS=(
--build-arg BASE_OS=${BASE_OS_URL}
--build-arg PACKAGENAME=$3
--build-arg NAME="${NAME}"
--build-arg PROCESS_NAME="${PROCESS_NAME}"
)

TARGET="$NAME:"$2
# Build the target and tag with the full tag.
docker build "${BUILD_ARGS[@]}" -f "dockerfile/Dockerfile-${BASE_OS}"  -t "$TARGET" $ROOT 



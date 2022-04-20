
FROM python:3

RUN pip3 install ansible==2.9.4

RUN pip3 install -i https://test.pypi.org/simple/ cotea==1.2

RUN apt update && apt install openssh-server sudo sshpass -y

ADD src/ /home/ubuntu

RUN chmod a+x /home/ubuntu/configure_ssh.sh

RUN chmod a+x /home/ubuntu/install_golang.sh

RUN /home/ubuntu/configure_ssh.sh

RUN /home/ubuntu/install_golang.sh

ENV ANSIBLE_HOST_KEY_CHECKING=False

ENTRYPOINT ["sleep", "infinity"]

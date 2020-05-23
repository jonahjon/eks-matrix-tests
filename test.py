
#!/usr/bin/env python3
import fire
import logging
import os
import subprocess
import sys
from pathlib import Path
from kubernetes import config

"""
This script serves as a *simplified* interface to e2e_test.go (which has too
many args and hard-to-find help text).
Like this Makefile which wraps various complicated scripts:
https://github.com/kubernetes/kubernetes/blob/master/build/root/Makefile
"""


def test(focus='',
         skip='',
         parallelism='',
         cluster_name='',
         args='',
         region=os.environ.get('AWS_DEFAULT_REGION', 'us-west-2'),
         kubeconfig=os.environ.get('KUBECONFIG',
                                   str(Path.home().joinpath('.kube/config')))):
    """Tests an EKS cluster
    Args:
      focus: Regexp that matches the tests to be run. Defaults to "".
      skip: Regexp that matches the tests that needs to be skipped. Defaults to "".
      parallelism: Number of ginkgo nodes to run. Defaults to ginkgo default
        (cores - 1).
      kubeconfig: Path to kubeconfig. Defaults to "$KUBECONFIG" or
        "$HOME/.kube/config".
    """
    if not cluster_name:
        cluster_name = guess_cluster_name(kubeconfig)
    if args:
        arguement=args
    else:
        args = 'ginkgo -p'.split()
        if focus:
            args.append(f'-focus={focus}')
        if skip:
            args.append(f'-skip={skip}')
        if parallelism:
            args.append(f'-nodes={parallelism}')
        
    args.append('./')
    args.append('--')
    if kubeconfig:
        args.append(f'-kubeconfig={kubeconfig}')

    print(' '.join(args))

    process = subprocess.run(
        args=arguement,
        cwd='./go',
        check=True,
        # aws-sdk-go expects AWS_REGION
        env={**os.environ, 'AWS_REGION': region},
    )


def guess_cluster_name(kubeconfig) -> str:
    try:
        _, current_context = config.list_kube_config_contexts(kubeconfig)
        cluster_arn = current_context['context']['cluster']
        cluster_name = cluster_arn.split('/')[1]
        return cluster_name
    except:
        return ''


if __name__ == "__main__":
    logging.basicConfig(stream=sys.stdout)
    fire.Fire(test)
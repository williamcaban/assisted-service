import os
import utils
import deployment_options


def main():
    deploy_options = deployment_options.load_deployment_options()

    utils.verify_build_directory(deploy_options.namespace)

    src_file = os.path.join(os.getcwd(), 'deploy/roles/default_role.yaml')
    dst_file = os.path.join(os.getcwd(), 'build', deploy_options.namespace, 'default_role.yaml')

    with open(src_file, "r") as src:
        with open(dst_file, "w+") as dst:
            data = src.read()
            data = data.replace('REPLACE_NAMESPACE', f'"{deploy_options.namespace}"')
            print("Deploying {}".format(dst_file))
            dst.write(data)

    utils.apply(
        target=deploy_options.target,
        namespace=deploy_options.namespace,
        profile=deploy_options.profile,
        file=dst_file
    )


if __name__ == "__main__":
    main()

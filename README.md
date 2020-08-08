# vault-auto-unsealer
Ported from Ruby. Thanks to https://github.com/InQuicker/vault-auto-unsealer

**vault-auto-unsealer** is an application to assist in running [Vault](https://www.vaultproject.io/) in a [Kubernetes](http://kubernetes.io/) cluster.
It runs a control loop that will unseal Vault if necessary, so that human intervention is not needed to unseal Vault if/when Vault's pod is restarted by Kubernetes.
Of course, this means that the value of Vault's unseal process is foregone, but this may be an acceptable compromise when running in an environment where the unseal key is protected via some other means, and the ability to keep Vault running without human intervention when it restarts is critical.

## Configuration

Required environment variables:

* `VAULT_ADDR`: The address of the Vault server to operate on. Example: http://127.0.0.1:8200
* `UNSEAL_KEY`: The raw decrypted unseal key.

## Usage

When first deployed, you dont need to set the environment variable `UNSEAL_KEY`.
vault-auto-unsealer will initialize Vault, configured for a single unseal key.
It will then print the unseal key and initial root token to standard output and will halt.

Then set the `UNSEAL_KEY` to the printed key and redeploy the container.

## Legal

vault-auto-unsealer is released under the MIT license.
See `LICENSE` for details.
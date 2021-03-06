//! This file is a single use binary that will allow you to register your validator ethereum key

#[macro_use]
extern crate serde_derive;
#[macro_use]
extern crate lazy_static;

use std::time::Duration;

use clarity::PrivateKey as EthPrivateKey;
use contact::client::Contact;
use cosmos_peggy::send::update_peggy_eth_address;
use deep_space::{coin::Coin, private_key::PrivateKey as CosmosPrivateKey};
use docopt::Docopt;
use rand::{thread_rng, Rng};
use url::Url;

#[derive(Debug, Deserialize)]
struct Args {
    flag_cosmos_phrase: String,
    flag_ethereum_key: Option<String>,
    flag_cosmos_rpc: String,
    flag_fees: String,
}

lazy_static! {
    pub static ref USAGE: String = format!(
        "Usage: {} --cosmos-phrase=<key> [--ethereum-key=<key>] --cosmos-rpc=<url> --fees=<denom>
        Options:
            -h --help                     Show this screen.
            --cosmos-phrase=<ckey>    The Cosmos private key of the validator. Must be saved when you generate your key
            --ethereum-key=<ekey>     (Optional) The Ethereum private key to register, will be generated if not provided
            --cosmos-rpc=<curl>       The Cosmos Legacy RPC url, usually the validator. This will need to be manually enabled
            --fees=<denom>            The Cosmos Denom in which to pay Cosmos chain fees
        About:
            Special purpose binary for bootstrapping Peggy chains. This will submit and optionally
            generate an Ethereum key that will be used to sign messages on behalf of your Validator
            on the Cosmos blockchain running the Peggy module. Be aware this Ethereum key must be kept
            safe as you can be slashed for losing it.
            Written By: {}
            Version {}",
        env!("CARGO_PKG_NAME"),
        env!("CARGO_PKG_AUTHORS"),
        env!("CARGO_PKG_VERSION"),
    );
}

const TIMEOUT: Duration = Duration::from_secs(60);

#[actix_rt::main]
async fn main() {
    env_logger::init();
    // On Linux static builds we need to probe ssl certs path to be able to
    // do TLS stuff.
    openssl_probe::init_ssl_cert_env_vars();

    let args: Args = Docopt::new(USAGE.as_str())
        .and_then(|d| d.deserialize())
        .unwrap_or_else(|e| e.exit());
    let cosmos_key = CosmosPrivateKey::from_phrase(&args.flag_cosmos_phrase, "")
        .expect("Failed to parse validator key");
    let ethereum_key = if let Some(key) = args.flag_ethereum_key {
        key.parse().expect("Invalid Ethereum Private key!")
    } else {
        let mut rng = thread_rng();
        let key: [u8; 32] = rng.gen();
        let key = EthPrivateKey::from_slice(&key).unwrap();
        println!(
            "No Ethereum key provided, your generated key is {} -> {}",
            key,
            key.to_public_key().unwrap()
        );
        key
    };
    let cosmos_url = Url::parse(&args.flag_cosmos_rpc).expect("Invalid Cosmos RPC url");
    let cosmos_url = cosmos_url.to_string();
    let cosmos_url = cosmos_url.trim_end_matches('/');
    let fee_denom = args.flag_fees;

    let contact = Contact::new(&cosmos_url, TIMEOUT);
    let fee = Coin {
        denom: fee_denom,
        amount: 1u64.into(),
    };

    update_peggy_eth_address(&contact, ethereum_key, cosmos_key, fee.clone())
        .await
        .expect("Failed to update Eth address");

    let eth_address = ethereum_key.to_public_key().unwrap();
    let cosmos_address = cosmos_key.to_public_key().unwrap().to_address();
    println!(
        "Registered Ethereum address {} for validator address {}",
        eth_address, cosmos_address
    )
}

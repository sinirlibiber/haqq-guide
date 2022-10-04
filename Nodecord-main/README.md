# Nodecord


Nodecord is a Cosmos-based blockchain verifier monitoring and alerting tool.

Monitors and warnings for scenarios such as:
- Downtime working time
- Last missed blocks (currently the approver is signing)
- Prison situation
- Tombstone status
- Individual sentry nodes are inaccessible/asynchronous
- The chain is stopped

Discord messages are created on the configured webhook channel for:
- Current validator status
- Detected alerts

## Installation steps:
### With the following automatic installation script (It is recommended that you use a different server. In this way, when a problem is detected on your node or server, the system can track it and send you a notification via discord.)

```
wget -O NODECORD.sh https://raw.githubusercontent.com/Nodeist/Nodecord/main/NODECORD && chmod +x NODECORD.sh && ./NODECORD.sh
```

## Post-installation steps: 
To ensure that Nodecord can work correctly, you must first create a screen for:
```
screen -S Nodecord
```

### Configuration settings:
Create a configuration file.
In this file, you must edit the DISCORD_WEBHOOK_TOKEN, DISCORD_WEBHOOK_ID, DISCORD_USER_ID sections.
Learn more about creating a Discord webhook [here](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks ) you can find.

After creating the webhook, copy your url. this will be something like:
`https://discord.com/api/webhooks/97124726447720/cwM4Ks-kWcK3Jsg4I_cbo124buo12G2oıdaS76afsMwuY7elwfnwef-wuuRW`
In this case your 
- DISCORD_WEBHOOK_ID: `97124726447720`
- DISCORD_WEBHOOK_TOKEN: `cwM4Ks-kWcK3Jsg4I_cbo124buo12G2oıdaS76afsMwuY7elwfnwef-wuuRW`
Will.
- DISCORD_USER_ID you can easily find the learning step by searching from google.


Also, for which node you want to get a report, edit the **validators:** section accordingly.
- Name: Network name (Kujira, Osmosis,Quicksilver, etc.)
- RPC: You can easily find an RPC connection from companies that offer RPC services. I usually follow Polkachu.
Example rpc for Kujira: **https://kujira-rpc.polkachu.com/**
- Address: This is neither your wallet address nor your valoper address. Be careful with this one. It should be your consensus address. You can find it in explorers.
- Chain-id: haqq_54211-2

Next up in the sentry section:
- name: The name of your server. you can give any name.
- grpc: ip address of your server + rpc port

If you are running the same node on more than one server. (if you are hosting a backup server)
- name
- grpc

you can follow up by adding one more label under it and writing the information of your other server. 
'nano~/nodecord/config.open your file by typing `yaml'. you will see a screen similar to the one below. save the necessary places by editing them.

```
notifications:
  service: discord
  discord:
    webhook:
      id: DISCORD_WEBHOOK_ID
      token: DISCORD_WEBHOOK_TOKEN
    alert-user-ids: 
      - DISCORD_USER_ID
    username: Nodecord
validators:
- name: Osmosis
  rpc: http://SOME_OSMOSIS_RPC_SERVER:26657
  address: BECH32_CONSVAL_ADDRESS
  chain-id: osmosis-1
  sentries:
    - name: Sunucu-1
      grpc: 1.2.3.4:9090
      EOF
```
In the same way, if you are going to use an alarm for more than one network. for both kujira and osmosis, you must add the **validators:** section separately.
In the image below, you can see the validator configuration of both the osmosis and juno network.

![nodeist](https://i.hizliresim.com/hplawtm.png )

### Starting the monitor

You can start the monitor with the following code:

"'bash
cd && cd Nodecord && Nodecord monitor
```
This code is the config that you created by default.it pulls data from the yaml file and starts tracking. 
I am using two separate yaml files for the testnet and mainnet. 
and a testnet for the testnet.you can create a yaml file. and you can run this separately with the code below.

"'bash
cd && cd Nodecord && Nodecord monitor -f~/testnet.yaml
```

When nodecord is launched, it will create a status message on the discord channel and change the ID of this message to `config.he will add to yaml.. Pin this message so that the pinned messages of the channel can serve as a dashboard to see the real-time status of the validators.

![Nodeist](https://i.hizliresim.com/6qt5b5t.png )

It will send warning messages when any error situation is detected.


![Nodeist](https://i.hizliresim.com/8ow2s04.png )

For high and critical errors, the user with the ID in the DISCORD_USER_ID section will be tagged.

![Nodeist](https://i.hizliresim.com/2g4vd1k.png )

It will send information message when the errors are fixed.


## Reference List
This project [Strangelove Ventures](https://github.com/strangelove-ventures ) was inspired by.

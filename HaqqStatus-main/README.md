# HaqqStatus

This system will alert you with telegram about jails, inactive status missed blocks and etc. Also it sends you every hour short info about your node status.

Instruction:

1. Create telegram bot via `@BotFather`, customize it and `get bot API token` ([how_to](https://www.siteguarding.com/en/how-to-get-telegram-bot-api-token)).
2. Create at least 2 groups: `alarm` and `log`. Customize them, add your bot into your chats and `get chats IDs` ([how_to](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id)).
3. Connect to your server and create `status` folder in the `$HOME directory` with `mkdir $HOME/status/`.
4. In this folder, `$HOME/status/`, you have to create `cosmos.sh` file with `nano $HOME/status/cosmos.sh`. You don't have to do any edits on `cosmos.sh` file, it's ready to use.
> You can find `cosmos.sh` in this repository.
5. In this folder, `$HOME/status/`, you have to create `haqq.conf` file with `nano $HOME/status/haqq.conf`. Customize it.
> You can find `haqq.conf` in this repository.
6. Install some packages with `sudo apt-get install jq sysstat bc smartmontools fdisk -y`.
7. Run `bash cosmos.sh` to check your settings. Normal output:
```
root@gumusbey:~/status# bash cosmos.sh
 
exp/me >>>>>> 247490/247490.
gap >>>>>>>>> 0 blocks.
chain >>>>>>> alive.
consensus >>> 0.00.
block_time >> 6.33 sec.
priv_key >>>> right.
_active >>>>> false.
gov >>>>>>>>> no unvoted proposals.

_upgrade >>>> v1.1.0.
_time_left >> 15h 18m.
_appr_time >> Sep 27, 16:36.

root@gumusbey:~/status#
```
9. Add some rules with `chmod u+x $HOME/status/cosmos.sh`.
10. Edit crontab with `crontab -e`.
```
# status
1,11,21,31,41,51 * * * * bash $HOME/status/cosmos.sh >> $HOME/status/cosmos.log 2>&1
```
11. Check your logs with `cat $HOME/status/cosmos.log` or `tail $HOME/status/cosmos.log -f`.


Referance: Status By cyberomanov


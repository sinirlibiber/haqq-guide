# Nodecord


Nodecord, Cosmos tabanlı bir blockchain doğrulayıcı izleme ve uyarı aracıdır.

Aşağıdakiler gibi senaryolar için izler ve uyarılar:
- Kesinti süresi çalışma süresi
- Son kaçırılan bloklar (şu anda onaylayıcı imzalıyor)
- Hapis durumu
- Mezar taşı durumu
- Bireysel nöbetçi düğümleri erişilemiyor/eşzamansız
- Zincir durduruldu

Aşağıdakiler için yapılandırılmış web kancası kanalında Discord mesajları oluşturulur:
- Geçerli doğrulayıcı durumu
- Algılanan uyarılar

## Kurulum adımları:
### Aşağıdaki otomatik kurulum scripti ile ( Önerilen farklı bir sunucu kullanmanızdır. Bu sayede node'unuzda ve ya sunucunuzda sorun tespit edildiğinde sistem takip edebilir ve size discord üzerinden bildirim gönderebilir.)

```
wget -O NODECORD.sh https://raw.githubusercontent.com/Nodeist/Nodecord/main/NODECORD && chmod +x NODECORD.sh && ./NODECORD.sh
```

## Kurulum sonrası adımlar: 
Nodecord'un doğru şekilde çalışabilmesini sağlamak için için öncelikle bir screen oluşturmalısınız:
```
screen -S Nodecord
```

### Konfigürasyon ayarları:
Bir konfigürasyon dosyası oluşturun.
Bu dosyada DISCORD_WEBHOOK_TOKEN, DISCOR_WEBHOOK_ID, DISCORD_USER_ID kısımlarını düzenlemelisiniz.
Discord webhook oluşturmak ile ilgili daha fazla bilgiyi [burada](https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks) bulabilirsiniz.

Webhook oluşturduktan sonra url'nizi kopyalayın. şuna benzeyecektir:
`https://discord.com/api/webhooks/97124726447720/cwM4Ks-kWcK3Jsg4I_cbo124buo12G2oıdaS76afsMwuY7elwfnwef-wuuRW`
Bu durumda sizin 
- DISCORD_WEBHOOK_ID: `97124726447720`
- DISCORD_WEBHOOK_TOKEN: `cwM4Ks-kWcK3Jsg4I_cbo124buo12G2oıdaS76afsMwuY7elwfnwef-wuuRW`
Olacaktır.
- DISCORD_USER_ID öğrenme adımını googledan araştırarak kolayca bulabilirsiniz.


Ayrıca hangi node için rapor almak istiyorsanız **validators:** bölümünü de ona göre düzenleyin. 
- Name: Network ismi (Kujira, Osmosis,Quicksilver vs.)
- RPC: RPC hizmeti sunan şirketlerden kolayca RPC bağlantısı bulabilirsiniz. Ben genellikle Polkachu'yu takip ediyorum. 
Kujira için örnek rpc: **https://kujira-rpc.polkachu.com/**
- Adress: Bu ne cüzdan adresiniz ne de valoper adresinizdir. Buna dikkat edin. Consensüs adresiniz olmalı. Explorer'lardan bulabilirsiniz.
- Chain-id: Kujira örneği için **Kaiyo-1**

Sıra sentry bölümünde:
- name: Sunucunuzun ismi. herhangi bir isim verebilirsiniz.
- grpc: Sunucunuzun ip adresi + rpc portu

Eğer aynı node birden fazla sunucuda çalıştırıyorsanız. (yedek bir sunucu barındırıyorsanız)
- name
- grpc

labellerinden altına birer tane daha ekleyip diğer sunucunuzun bilgilerini de yazarak takip edebilirsiniz. 
`nano ~/Nodecord/config.yaml` yazarak dosyanızı açın. aşağıdakine benzer bir ekran göreceksiniz. gerekli yerleri düzenleyerek kaydedin.

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
Yine aynı şekilde birden fazla ağ için alarm kullanacaksanız. hem kujira hem osmosis için ayrı ayrı **validators:** bölümünü komple eklemelisiniz.
Aşağıdaki resimde hem osmosis hem juno ağının validatör yapılandırmasını görebilirsiniz.

![nodeist](https://i.hizliresim.com/hplawtm.png)

### Monitörü başlatma

Aşağıdaki kod ile monitörü başlatabilirsiniz:

```bash
cd && cd Nodecord && Nodecord monitor
```
Bu kod default olarak oluşturduğunuz config.yaml dosyasından verileri çeker ve takibe başlar. 
Ben testnet ve mainnet için iki ayrı yaml dosyası kullanıyorum. 
testnet için bir de testnet.yaml dosyası oluşturabilirsiniz. ve bunu ayrı screen de aşağıdaki kod ile çalıştırabilirsiniz.

```bash
cd && cd Nodecord && Nodecord monitor -f ~/testnet.yaml
```

Nodecord başlatıldığında, discord kanalında bir durum mesajı oluşturacak ve bu mesajın ID'sini `config.yaml`a ekleyecektir.. Bu mesajı sabitleyin, böylece kanalın sabitlenmiş mesajları, doğrulayıcıların gerçek zamanlı durumunu görmek için bir gösterge panosu görevi görebilir.

![Nodeist](https://i.hizliresim.com/6qt5b5t.png)

Herhangi bir hata durumu tespit edildiğinde uyarı mesajları gönderecektir.


![Nodeist](https://i.hizliresim.com/8ow2s04.png)

Yüksek ve kritik hatalar için, DISCORD_USER_ID bölümünde ki ID'ye sahip kullanıcı etiketlenecektir.

![Nodeist](https://i.hizliresim.com/2g4vd1k.png)

Hatalar giderildiğinde bilgi mesajı gönderecektir.


## Referans Listesi
Bu proje [Strangelove Ventures](https://github.com/strangelove-ventures)'den esinlenilmiştir.

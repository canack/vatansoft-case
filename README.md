# VatanSoft Case

## VatanSoft tarafından gönderilen case için yazılan kodları ve açıklamaları içerir.

#### Verilen case için şunlar yapıldı:
- Todo-api servisi yazıldı
- Echo framework kullanıldı
- Gorm kullanıldı
- Veri tabanı olarak mySQL kullanıldı 
- Tutarlı ve izole bir ortam için container hale getirildiler
- Servisler arası iletişim için docker-compose kullanıldı
- API kullanım talimatları için swagger kullanıldı

## Kullanım:

### Bu case'in kaynak kodlarını; docker&docker-swarm kullanarak test edebilirsiniz.

###  Docker:
```shell
# Bu talimatları uygulamadan önce, docker ve docker-compose aracının sisteminizde yüklü olduğundan emin olun.

git clone https://github.com/canack/vatansoft-case.git
cd vatansoft-case
docker-compose up # Gerekli tüm servisleri başlatır.

# Servislerin başlatılma süresi, makina gücüne göre biraz zaman alabilir.
# Echo framework'un başlatılıp, istek almaya hazır hale geldiğini gördüğünüzde deneyebilirsiniz. 
# İşlemler tamamlandıktan sonra localhost:1323/swagger/index.html adresi üzerinden test edebilirsiniz

# Kapatmak için:
# CTRL + C
docker-compose down
```


---


###### Gönderilen case'i en temiz ve basit haliyle ortaya çıkartmayı amaçladım. Sadeliği korumak amacıyla derinlemesine kodlama yapmaktan kaçındım.

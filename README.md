# Agent Allocation

### Pada backend product ini, diterapkan beberapa fitur agent allocation, yang mana antara lain:
### 1. Fitur agent berstatus aktif melayani maksimal 2 customer, dimana apabila tiap agent aktif telah di alokasikan kepada 2 customer dengan status chat aktif, maka channel baru akan di alokasikan kepada agent aktif dengan jumlah chat aktif paling sedikit. Apabila tidak terdapat agent yang aktif, secara otomatis agent akan di assign kepada customer dengan menghitung jumlah chat aktif paling sedikit dari setiap agent.
### 2. Fitur customer menginisiasi channel. Selama chat masih aktif, customer tidak dapat menginisiasi channel baru.
### 3. Fitur agent me-resolve channel.
### 4. Fitur messaging antar customer dan agent.
### 5. Fitur Login dan Logout untuk masing-masing agent.
### 6. Fitur Agent dapat melihat channel aktif yang dialokasikan kepada agent tersebut.
### 7. Fitur Agent dan Customer dapat melihat pesan sebelumnya di channel tersebut.

## Get Started
### 1. LOGIN 
Daftar agent dan customer telah tersedia dan dapat dipakai, antara lain sebagai berikut:
1. Tabel agent

| id | username | password |
| --- | --- | --- |
| 1 | annasianna | strongpass1 |
| 2 | rosesimawar | strongpass2 |
| 3 | biruisblue 3 | strongpass3 |

2. Tabel client

| id | username | password |
| --- | --- | --- |
| 1 | daungugur | strongpass1 |
| 2 | mawarberduri | strongpass2 |
| 3 | langitbiru | strongpass3 |
| 4 | jalankenangan | strongpass4 |
| 5 | sepaturoda | strongpass5 |
| 6 | tumblrbiru | strongpass6 |
| 7 | jejakkaki | strongpass7 |
| 8 | tamanbunga | strongpass8 |

Pada postman, dapat digunakan method `POST` dengan endpoint :

```
ip_address:8080/agent/login
```
atau:
```
ip_address:8080/customer/login
```
lalu pada body pilih form value, kemudian dapat diisi dengan format sebagai berikut:
| key | value |
| --- | --- |
| username | tumblrbiru |
| password | strongpass6 |

Apabila proses berhasil, maka akan muncul data diri dari customer ataupun agent beserta token yang dapat digunakan untuk proses selanjutnya (diinput pada Authorization : bearer token).

### 2. Next Step
Langkah berikutnya setelah login berhasil dilakukan, agent ataupun customer dapat mengakses endpoint-endpoint berikut:
| No. | Method | Role | Endpoint | Form Value input key | Need Token? | Endpoint Example | Function |
| --- | --- | --- | --- | --- | --- | --- | --- |
| 1 | GET | Agent | ip_address/agent/:agent_id/chat/active | - | yes | ip_address/agent/1/chat/active | Agent can see active channel |
| 2 | GET | Agent | ip_address/agent/:agent_id/chat | channel_id, customer_id, agent_id | yes | ip_address/agent/1/chat | Agent can see messages from that channel |
| 3 | POST | Agent | ip_address/agent/:agent_id/chat/send | RecipientID, text_message | yes | ip_address/agent/1/chat/send | Agent can send messages to customer |
| 4 | POST | Agent | ip_address/agent/:agent_id/chat/resolve | customer_id | yes | ip_address/agent/1/chat/resolve | Agent can resolve a channel |
| 5 | PUT | Agent | ip_address/agent/:agent_id/logout | - | yes | ip_address/agent/1/logout | Agent can logout |
| 6 | POST | Customer | ip_address/customer/:customer_id/chat/initiate | - | yes | ip_address/customer/1/chat/initiate | Customer can initiate a channel to chat with agent. |
| 7 | GET | Customer | ip_address/customer/:customer_id/chat | channel_id | yes | ip_address/customer/1/chat | Customer can see chat in active channel. |
| 8 | POST | Customer | ip_address/customer/:customer_id/chat/send | text_message | yes | ip_address/customer/1/chat/send | Customer can send message to agent. |
| 9 | PUT | Customer | ip_address/customer/:customer_id/logout | - | yes | ip_address/customer/1/logout | Customer can logout |

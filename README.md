# RabbitMQ Fair Dispatch Project

โปรเจกต์นี้ประกอบด้วย 2 ส่วนหลัก:
- go-publish: สำหรับ publish ข้อมูลเข้า RabbitMQ
- go-consume: สำหรับ consume ข้อมูลจาก RabbitMQ

## โครงสร้างโปรเจกต์

```
go-consume/
  main.go
  ...
go-publish/
  main.go
  ...
rabbitmq/
  Dockerfile
```

## วิธีการใช้งาน

### 1. ติดตั้ง RabbitMQ ด้วย Docker

ไปที่โฟลเดอร์ `rabbitmq` และรันคำสั่ง:

```sh
docker build -t my-rabbitmq .
docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 my-rabbitmq
```

- RabbitMQ Management UI: http://localhost:15672 (user: rabbituser, pass: rabbitpassword)


### 2. รัน go-consume (Consumer) ก่อน

เปิดเทอร์มินัลใหม่ แล้วรัน:

```sh
cd go-consume
go run main.go
```

หากต้องการทดสอบการ fair dispatch ให้เปิดหลายเทอร์มินัล แล้วรัน consumer หลายตัวพร้อมกัน เช่น เปิด 2-3 เทอร์มินัลแล้วรันคำสั่งข้างต้นในแต่ละหน้าต่าง


### 3. รัน go-publish (Publisher)

```sh
cd go-publish
go run main.go
```

#### ส่งข้อความแบบ custom
สามารถระบุข้อความที่ต้องการส่งเองได้ เช่น

```sh
go run main.go "ข้อความที่ต้องการส่ง"
```
เช่น
```sh
go run main.go "Hello RabbitMQ!"
```

- ตรวจสอบให้แน่ใจว่า RabbitMQ ทำงานอยู่ก่อนรัน consumer/publisher
- สามารถปรับแต่งโค้ดใน `main.go` ทั้งสองฝั่งได้ตามต้องการ



# راهنمای تنظیمات پروژه Ambridge

## فایل تنظیمات

پروژه Ambridge از یک فایل تنظیمات مرکزی استفاده می‌کند که در `config/env.go` قرار دارد. این فایل تمام تنظیمات مورد نیاز برای اجرای پروژه را مدیریت می‌کند.

## متغیرهای محیطی

تنظیمات از طریق متغیرهای محیطی یا فایل `.env` قابل تنظیم هستند. یک نمونه از فایل `.env` در `env.sample` موجود است. برای استفاده، آن را به `.env` کپی کنید و مقادیر را متناسب با محیط خود تغییر دهید.

### متغیرهای پایگاه داده

```
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=your_password
MYSQL_DATABASE=ambridge
```

### متغیرهای JWT

```
JWT_SECRET=your_secret_key
JWT_EXPIRATION=24
```

### متغیرهای سرور

```
SERVER_PORT=8080
```

## استفاده از تنظیمات در کد

برای استفاده از تنظیمات در کد، ابتدا پکیج `config` را import کنید:

```go
import "ambridge-backend/config"
```

سپس می‌توانید از توابع دسترسی به تنظیمات استفاده کنید:

### تنظیمات پایگاه داده

```go
// دریافت رشته اتصال به پایگاه داده
dsn := config.GetMySQLDSN()

// یا دسترسی به تنظیمات جداگانه
host := config.GetMySQLHost()
port := config.GetMySQLPort()
user := config.GetMySQLUser()
password := config.GetMySQLPassword()
database := config.GetMySQLDatabase()
```

### تنظیمات JWT

```go
secret := config.GetJWTSecret()
expiration := config.GetJWTExpiration() // مقدار به ساعت
```

### تنظیمات سرور

```go
port := config.GetServerPort()
```

## نکات مهم

1. قبل از استفاده از تنظیمات، باید تابع `config.LoadConfig()` را فراخوانی کنید. این کار در `database.InitDB()` انجام می‌شود.
2. اگر متغیر محیطی تنظیم نشده باشد، از مقدار پیش‌فرض استفاده می‌شود.
3. برای امنیت بیشتر، مقادیر حساس مانند رمز عبور و کلید JWT را در فایل `.env` قرار دهید و آن را در مخزن گیت قرار ندهید.
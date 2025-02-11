
Giải thích:
1. logConfig Struct:
- **LogLevel**: Đặt mức độ log (debug, info, warn, error).
- **LogType**: Chọn loại log (FILE hoặc DEFAULT).
- **LogDir**: Thư mục lưu trữ log.
- **MaxSize**: Kích thước tối đa của file log (MB).
- **MaxBackup**: Số lượng file log backup tối đa.
- **MaxAge**: Số ngày lưu trữ log.

2. setAppLogger Function:
- Thiết lập formatter và mức độ log.
- Nếu **LogType** là **FILE**, nó sẽ tạo một file log mới và thiết lập đầu ra log là cả os.Stdout và file log.

3. createNewLogFile Function:
- Tạo thư mục log nếu chưa tồn tại.
- Tạo một file log mới với tên chứa ngày hiện tại.
- Sử dụng **lumberjack.Logger** để quản lý việc xoay vòng log dựa trên kích thước file, số lượng backup, và tuổi thọ của file.

Cách hoạt động:
- Log Rotation: Khi file log đạt đến kích thước tối đa **MaxSize**, lumberjack sẽ tự động tạo một file log mới và đổi tên file cũ thành LogFile-<date>.log.gz, v.v.

- MaxBackup: Giới hạn số lượng file log backup (ở đây là 3).

- MaxAge: File log cũ hơn **MaxAge** ngày sẽ bị xóa tự động.

Lưu ý:
- Bạn có thể điều chỉnh các tham số như MaxSize, MaxBackup, và MaxAge để phù hợp với nhu cầu của bạn.
- Đảm bảo rằng thư mục log (LogDir) có quyền ghi để chương trình có thể tạo và ghi file log.
- Với cách triển khai này, bạn sẽ có một hệ thống quản lý log hoàn chỉnh, tự động xoay vòng log dựa trên kích thước và thời gian, đồng thời quản lý số lượng file log backup một cách hiệu quả.

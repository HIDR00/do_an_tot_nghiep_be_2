# Mono-Base Project

## Giới thiệu
Mono-Base là một dự án mẫu sử dụng GoLang, được thiết kế để làm cơ sở cho các ứng dụng lớn hơn, theo kiến trúc microservice hoặc monolith đơn giản. Dự án này được tổ chức theo cấu trúc thư mục tiêu chuẩn của Go, bao gồm việc tách biệt rõ ràng giữa các thành phần internal, external và configurations.

## Cấu trúc thư mục

- **cmd/**: Chứa các điểm nhập của ứng dụng, bao gồm các dịch vụ CLI và HTTP.
    - **cli/**: Các tập lệnh dòng lệnh.
    - **http/**: Điểm khởi động cho các dịch vụ web.
- **configs/**: Cấu hình của dự án, bao gồm mẫu và file cấu hình chính.
- **internal/**: Thư mục cho các định nghĩa và logic nội bộ của dự án.
    - **entities/**: Định nghĩa các mô hình dữ liệu cốt lõi.
    - **external/**: Interfaces để tương tác với các dịch vụ bên ngoài.
    - **infrastructure/**: Cài đặt cơ sở hạ tầng như cơ sở dữ liệu và bộ nhớ cache.
    - **mocks/**: Định nghĩa mocks để hỗ trợ kiểm thử.
    - **repositories/**: Cài đặt các repositories để truy cập và quản lý dữ liệu.
    - **services/**: Cài đặt logic xử lý nghiệp vụ.
    - **usecases/**: Logic nghiệp vụ cụ thể.
- **pkg/**: Thư mục chứa các packages có thể tái sử dụng ở tầm rộng hơn.
    - **config/**: Các tiện ích đọc cấu hình.
    - **constants/**: Định nghĩa các hằng số toàn cục.
    - **error/**: Xử lý lỗi tùy chỉnh.
    - **logger/**: Cài đặt logging.
    - **psql-helper/**: Tiện ích hỗ trợ cho PostgreSQL.
    - **utils/**: Các hàm tiện ích khác.
- **tools/**: Các công cụ hỗ trợ cho việc phát triển, như docker-compose.
- **volumes/**: Thư mục để chứa dữ liệu khi phát triển với Docker.

## Cài đặt và Khởi chạy

Để cài đặt và khởi chạy Mono-Base, hãy làm theo các bước sau:
1. Cài đặt các phụ thuộc:
   ```
   go mod tidy
   ```
2. Cấu hình ứng dụng:
   ```
   cp configs/config.yaml.sample configs/config.yaml
   # Sửa đổi config.yaml để phù hợp với môi trường của bạn
   ```
3. Cài đặt PostgreSQL và Redis bằng Docker:
   ```
   docker-compose -f tools/docker-compose.local.yml up -d
   ```
4. Chạy các migrations: (muốn chạy cái này thì phải mở docker và chạy hết tất cả container)
   ```
   go run mono-base/cmd/cli/migration
   ```
5. Khởi chạy ứng dụng:
   ```
   go run mono-base/cmd/http 
   ```
6. Tạo bản build chính thức cho ứng dụng
   ```
   go build mono-base/cmd/http
   ```
7. gen file wire_gen.go: https://www.notion.so/Wire-1d60b93656488053981ef2a57cf7c9cc  

## Contributing

Chúng tôi hoan nghênh mọi đóng góp từ cộng đồng. Nếu bạn muốn cải thiện Mono-Base, xin vui lòng tạo pull request hoặc gửi issue.

## License

Mono-Base được cấp phép dưới [MIT License](LICENSE).

---# Do_an_tot_nghiep_be

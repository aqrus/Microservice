import React from "react";

import { Container, Header, Main, Footer, Cards } from "@components";

const Home: React.FC = () => {
  return (
    <Container>
      <Header />
      User-service: Golang with Gin and mongodb: quản lý thông tin tài khoản người dùng, đăng ký, đăng nhập, quên mật khẩu, quản lý thông tin cá nhân

Cart-service: Golang with Gin and redis: quản lý giỏ hàng của người dùng, thêm sản phẩm, xoá sản phẩm, chuyển sản phẩm giữa các trang

Order-service: Golang with Gin and mongodb: quản lý đơn hàng, xem trạng thái đơn hàng, thanh toán đơn hàng, đặt hàng

Về việc tách riêng lẽ các service, nếu bạn muốn nâng cao hiệu suất và quản lý được tốt hơn, tôi đề nghị bạn tách chúng thành các microservice riêng biệt, mỗi service có một nhiệm vụ nhất định và tương tác với các service khác thông qua các HTTP API.
      <Main />
      <Cards />
      <Footer />
    </Container>
  );
};

export default Home;

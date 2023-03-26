import React from "react";
import Button from '@mui/material/Button';
import { Header, Main, Cards, Footer } from "components/css";

const Home: React.FC = () => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        minHeight: "100vh",
      }}
    >
      <Header />
      <Button variant="contained">Hello World</Button>
      <Main />
      <Cards />
      <Footer />
    </div>
  );
};

export default Home;

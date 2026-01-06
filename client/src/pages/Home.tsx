import React, { useState } from "react";
import Typewriter from "typewriter-effect";
import { Link } from "react-router-dom";
import { Button } from "@mui/material";

const Home: React.FC = () => {
    const [isShowButton, setIsShowButton] = useState(false);

    const hideButton = () => {
        setIsShowButton(false);
    };

    const showButton = () => {
        setIsShowButton(true);
    };

    return (
        <div
            style={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                height: "100vh",
                textAlign: "center",
                gap: "6rem",
                backgroundImage: "url('/images/home_background.jpg')",
                backgroundSize: "cover",
                backgroundPosition: "center",
                backgroundRepeat: "no-repeat",
            }}
        >
            <h1
                style={{
                    whiteSpace: "nowrap",
                    fontSize: "4rem",
                    margin: 0,
                    background: "linear-gradient(45deg, #2E3192, #1BFFFF)",
                    WebkitBackgroundClip: "text",
                    WebkitTextFillColor: "transparent",
                    backgroundClip: "text",
                }}
            >
                <Typewriter
                    onInit={(typewriter) => {
                        hideButton();
                        typewriter
                            .changeDelay(80)
                            .pauseFor(500)
                            .typeString("Welcome To CVWO Web Forum")
                            .callFunction(showButton)
                            .start();
                    }}
                />
            </h1>

            <Button
                variant="contained"
                color="primary"
                component={Link}
                to="/login"
                style={{
                    fontSize: "1.5rem",
                    padding: "1rem 2.5rem",
                    opacity: isShowButton ? 1 : 0,
                    transition: "opacity 1.5s ease-in-out",
                }}
            >
                Get Started
            </Button>
        </div>
    );
};

export default Home;

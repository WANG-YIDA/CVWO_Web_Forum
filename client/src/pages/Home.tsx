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
            }}
        >
            <style>
                {`
                    @keyframes gradient-flow {
                        0% { background-position: 0% 50%; }
                        50% { background-position: 100% 50%; }
                        100% { background-position: 0% 50%; }
                    }
                    @keyframes fade-in-up {
                        from { opacity: 0; transform: translateY(20px); }
                        to { opacity: 1; transform: translateY(0); }
                    }
                `}
            </style>
            <h1
                style={{
                    whiteSpace: "nowrap",
                    fontSize: "4rem",
                    margin: 0,
                    background: "linear-gradient(45deg, #2E3192, #1BFFFF, #2E3192)",
                    backgroundSize: "200% auto",
                    WebkitBackgroundClip: "text",
                    WebkitTextFillColor: "transparent",
                    backgroundClip: "text",
                    filter: "drop-shadow(2px 2px 4px rgba(0, 0, 0, 0.3))",
                    animation: "gradient-flow 5s ease 0s 2, fade-in-up 1s ease-out forwards",
                }}
            >
                <Typewriter
                    onInit={(typewriter) => {
                        hideButton();
                        typewriter
                            .changeDelay(40)
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

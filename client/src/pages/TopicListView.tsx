import TopicList from "../components/TopicList";
import AddIcon from "@mui/icons-material/Add";
import { Box, IconButton, keyframes, Typography } from "@mui/material";
import React from "react";

const fadeIn = keyframes`
    from { opacity: 0; transform: translateY(12px); }
    to { opacity: 1; transform: translateY(0); }
`;

const TopicListView: React.FC = () => {
    return (
        <div
            style={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                minHeight: "100vh",
                textAlign: "center",
                gap: "3rem",
                paddingTop: "2rem",
            }}
        >
            <Box sx={{ width: "100%", maxWidth: 800, padding: "0 16px", margin: "0 auto" }}>
                <Typography
                    component="h4"
                    sx={{
                        whiteSpace: "nowrap",
                        fontSize: { xs: "2.5rem", sm: "3rem", md: "3.5rem" },
                        mb: "3rem",
                        background: "linear-gradient(45deg, #2E3192, #1BFFFF)",
                        WebkitBackgroundClip: "text",
                        WebkitTextFillColor: "transparent",
                        backgroundClip: "text",
                        filter: "drop-shadow(2px 2px 4px rgba(0, 0, 0, 0.3))",
                        opacity: 0,
                        animation: `${fadeIn} 1.2s ease forwards`,
                    }}
                >
                    <strong>Discover Your Conversations</strong>
                </Typography>

                <Box sx={{ display: "flex", justifyContent: "flex-end", mr: 1, mb: -1 }}>
                    <IconButton
                        aria-label="add"
                        sx={{
                            width: 40,
                            height: 25,
                            borderRadius: "8px",
                            backgroundColor: "primary.main",
                            color: "#fff",
                            "&:hover": { backgroundColor: "primary.dark" },
                        }}
                    >
                        <AddIcon sx={{ fontSize: 16 }} />
                    </IconButton>
                </Box>
                <TopicList />
            </Box>
        </div>
    );
};

export default TopicListView;

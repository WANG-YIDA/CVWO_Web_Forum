import TopicList from "../components/TopicList";
import Topic from "../types/Topic";
import AddIcon from "@mui/icons-material/Add";
import { Box, IconButton, keyframes, Typography } from "@mui/material";
import React, { useEffect, useState } from "react";

const fadeIn = keyframes`
    from { opacity: 0; transform: translateY(12px); }
    to { opacity: 1; transform: translateY(0); }
`;

const TopicListView: React.FC = () => {
    const [topics, setTopics] = useState<Topic[]>([]);
    const [loading, setLoading] = useState<boolean>(true);

    useEffect(() => {
        const fetchTopics = async () => {
            try {
                // Request to View Topics
                const response = await fetch("http://localhost:8080/topics", { method: "GET" });
                const data_json = await response.json();

                //process data to get topics
                if (data_json.success && data_json.payload?.data) {
                    const topicListResult = JSON.parse(data_json.payload.data);

                    if (topicListResult.success && topicListResult.topics) {
                        const topicList: Topic[] = topicListResult.topics.map((topic: any) => ({
                            id: topic.id,
                            name: topic.name,
                            description: topic.description,
                            author: topic.author,
                            timestamp: new Date(topic.created_at),
                        }));

                        setTopics(topicList);
                    }
                } else {
                    console.error("Failed to GET topics: %s", data_json.error);
                }
            } catch (error) {
                console.error("Error fetching topics:", error);
            } finally {
                setLoading(false);
            }
        };

        fetchTopics();
    }, []);

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

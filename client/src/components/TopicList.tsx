import TopicItem from "./TopicItem";
import Topic from "../types/Topic";

import React from "react";
import { Box } from "@mui/material";

const TopicList: React.FC = () => {
    const topics: Topic[] = [
        {
            name: "topic_test1",
            description:
                "Test topic 1LLorem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!orem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!",
            author: "LucasW",
            timestamp: new Date(),
        },
        {
            name: "topic_test2",
            description: "Test topic 2",
            author: "LucasW",
            timestamp: new Date(),
        },
        {
            name: "topic_test3",
            description: "Test topic 3",
            author: "LucasW",
            timestamp: new Date(),
        },
    ];

    return (
        <Box sx={{ padding: 2 }}>
            {topics.map((topic) => (
                <TopicItem topic={topic} key={topic.name} />
            ))}
        </Box>
    );
};

export default TopicList;

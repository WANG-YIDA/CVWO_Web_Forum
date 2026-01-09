import TopicItem from "./TopicItem";
import Topic from "../types/Topic";

import React from "react";
import { Box } from "@mui/material";

interface Props {
    topics: Topic[];
}

const TopicList: React.FC<Props> = ({ topics }) => {
    return (
        <Box sx={{ padding: 2 }}>
            {topics.map((topic) => (
                <TopicItem topic={topic} key={topic.name} />
            ))}
        </Box>
    );
};

export default TopicList;

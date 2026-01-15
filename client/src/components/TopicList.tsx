import TopicItem from "./TopicItem";
import Topic from "../types/Topic";

import React from "react";
import { Box } from "@mui/material";

interface Props {
    topics: Topic[];
    user_id: number | null;
    onDeleteTopic: (topic_id: number) => void;
}

const TopicList: React.FC<Props> = ({ topics, user_id, onDeleteTopic }) => {
    if (!user_id) {
        return;
    }

    return (
        <Box sx={{ padding: 2 }}>
            {topics.map((topic) => (
                <TopicItem topic={topic} user_id={user_id} onDeleteTopic={onDeleteTopic} key={topic.id} />
            ))}
        </Box>
    );
};

export default TopicList;

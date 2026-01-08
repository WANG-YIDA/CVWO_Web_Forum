import PostItemPreview from "./PostItemPreview";
import Post from "../types/Post";
import React from "react";
import { Box, Typography } from "@mui/material";

const TopicList: React.FC = () => {
    const topic_name = "test";
    const posts: Post[] = [
        {
            topic: "test",
            title: "post_test1",
            body: "Test topic 1",
            author: "LucasW",
            timestamp: new Date(),
        },
        {
            topic: "test",
            title: "post_test2",
            body: "Lorem ipsLLorem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!orem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!um dolor sit amet consectetur adipisicing elit. Underword eligendi aut aliquam minus expedita repellendus magnam, quasi, iure omnis laboriosam quibusdam, corporis quam illo soluta nemo doloribus eius consequatur dignissimos.Lorem ipsum dolor sit amet consectetur adipisicing elit. Unde eligendi aut aliquam minus expedita repellendus magnam, quasi, iure omnis laboriosam quibusdam, corporis quam illo soluta nemo doloribus eius consequatur dignissimos.",
            author: "LucasW",
            timestamp: new Date(),
        },
        {
            topic: "test",
            title: "post_test3",
            body: "Test topic 3",
            author: "LucasW",
            timestamp: new Date(),
        },
    ];

    return (
        <Box sx={{ maxWidth: 1000, mx: "auto", px: 2 }}>
            <Typography fontSize={15} color="textSecondary" gutterBottom align="left">
                {"Topics > "}
                {topic_name}
            </Typography>
            <Box sx={{ mt: 2 }}>
                {posts.map((post) => (
                    <PostItemPreview post={post} key="" />
                ))}
            </Box>
        </Box>
    );
};

export default TopicList;

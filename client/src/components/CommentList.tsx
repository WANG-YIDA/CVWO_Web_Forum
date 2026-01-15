import CommentItem from "./CommentItem";
import Comment from "../types/Comment";

import React from "react";
import { Box, Typography } from "@mui/material";

const BasicCommentList: React.FC = () => {
    const comments: Comment[] = [
        {
            id: 1,
            user_id: 1,
            body:
                "Any fool can write code that a computer can understand.\n" +
                "Good programmers write code that humans can understand.\n" +
                " ~ Martin Fowler",
            author: "Benedict",
            timestamp: new Date(2022, 10, 28, 10, 33, 30),
        },
        {
            id: 1,
            user_id: 1,
            body: "Code reuse is the Holy Grail of Software Engineering.\n" + " ~ Douglas Crockford",
            author: "Casey",
            timestamp: new Date(2022, 11, 1, 11, 11, 11),
        },
        {
            id: 1,
            user_id: 1,
            body: "LLorem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!orem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!",
            author: "Duuet",
            timestamp: new Date(2022, 11, 2, 10, 30, 0),
        },
    ];

    return (
        <Box sx={{ padding: 2, maxWidth: 1000, margin: "0 auto" }}>
            <Typography variant="h5" color="textSecondary" gutterBottom align="left" marginBottom={2}>
                <strong>Comments:</strong>
            </Typography>
            {comments.map((comment) => (
                <CommentItem comment={comment} key="" />
            ))}
        </Box>
    );
};

export default BasicCommentList;

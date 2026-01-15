import PostItemPreview from "./PostItemPreview";
import Post from "../types/Post";
import React from "react";
import { Box, Divider, Link, Typography } from "@mui/material";

interface Props {
    posts: Post[];
    topicID: string;
    topicName: string;
}

const PostList: React.FC<Props> = ({ posts, topicID, topicName }) => {
    return (
        <Box sx={{ maxWidth: 1000, mx: "auto", px: 2 }}>
            <Typography fontSize={15} color="textSecondary" align="left">
                <Link href="/topics">Topics</Link>
                {" > "}
                {topicName}
            </Typography>

            <Divider sx={{ borderColor: "#bababeff", borderBottomWidth: 2 }} />

            <Box sx={{ mt: 2 }}>
                {posts.map((post) => (
                    <PostItemPreview post={post} topicID={topicID} key={post.id} />
                ))}
            </Box>
        </Box>
    );
};

export default PostList;

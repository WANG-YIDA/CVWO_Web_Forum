import Post from "../types/Post";
import CommentList from "../components/CommentList";
import PostItem from "../components/PostItem";
import AddIcon from "@mui/icons-material/Add";
import React from "react";
import { Box, IconButton } from "@mui/material";

const PostItemView: React.FC = () => {
    const post: Post = {
        topic: "test",
        title: "post_test2",
        body: "Lorem ipsLLorem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!orem ipsum dolor sit amet consectetur, adipisicing elit. At aliquam neque incidunt ratione enim minus, asperiores, labore earum totam numquam excepturi. Quod cupiditate amet quos repellat incidunt voluptatibus ad? Totam!um dolor sit amet consectetur adipisicing elit. Underword eligendi aut aliquam minus expedita repellendus magnam, quasi, iure omnis laboriosam quibusdam, corporis quam illo soluta nemo doloribus eius consequatur dignissimos.Lorem ipsum dolor sit amet consectetur adipisicing elit. Unde eligendi aut aliquam minus expedita repellendus magnam, quasi, iure omnis laboriosam quibusdam, corporis quam illo soluta nemo doloribus eius consequatur dignissimos.",
        author: "LucasW",
        timestamp: new Date(),
    };
    return (
        <div
            style={{
                display: "flex",
                flexDirection: "column",
                justifyContent: "center",
                alignItems: "center",
                minHeight: "100vh",
                textAlign: "center",
                gap: "1rem",
                paddingTop: "3rem",
            }}
        >
            <PostItem post={post} />
            <Box sx={{ display: "flex", justifyContent: "flex-end", mr: -95, mb: -9 }}>
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
            <CommentList />
        </div>
    );
};

export default PostItemView;

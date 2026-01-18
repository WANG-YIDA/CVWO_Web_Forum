import CommentItem from "./CommentItem";
import Comment from "../types/Comment";
import React from "react";
import { Box } from "@mui/material";

interface Props {
    comments: Comment[];
    user_id: number | null;
    topic_id: string | null;
    onDeleteComment: (comment_id: number) => void;
}

const CommentList: React.FC<Props> = ({ comments, user_id, topic_id, onDeleteComment }) => {
    if (!user_id || !topic_id) {
        return;
    }

    return (
        <Box sx={{ padding: 2, maxWidth: 750, width: "100%", margin: "0 auto" }}>
            {comments.map((comment) => (
                <CommentItem
                    comment={comment}
                    user_id={user_id}
                    topic_id={topic_id}
                    onDeleteComment={onDeleteComment}
                    key={comment.id}
                />
            ))}
        </Box>
    );
};

export default CommentList;

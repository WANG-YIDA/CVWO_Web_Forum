import Post from "../types/Post";
import { makeStyles } from "@mui/styles";
import { Box, Card, CardActions, CardContent, IconButton, Typography } from "@mui/material";
import ModeEditIcon from "@mui/icons-material/ModeEdit";
import DeleteIcon from "@mui/icons-material/Delete";
import React from "react";

type Props = {
    post: Post;
};

const useStyles = makeStyles(() => ({
    postTitle: {
        fontSize: 32,
        whiteSpace: "pre-wrap",
        paddingBottom: "0.5em",
        textAlign: "left",
        fontWeight: 500,
    },
    postBody: {
        fontSize: 16,
        textAlign: "left",
        lineHeight: 1.6,
        letterSpacing: "0.01em",
    },
    postCard: {
        maxWidth: 800,
        width: "100%",
        marginBottom: "3em",
        padding: "8px",
    },
    metadata: {
        fontSize: "12px !important",
        marginLeft: "auto",
    },
    cardActions: {
        display: "flex",
        justifyContent: "space-between",
        alignItems: "center",
    },
}));

const PostItem: React.FC<Props> = ({ post }) => {
    const classes = useStyles();

    const topic_name = "test";
    return (
        <Box sx={{ maxWidth: 1000, mx: "auto", px: 2 }}>
            <Typography fontSize={15} color="textSecondary" gutterBottom align="left" marginBottom={2}>
                {"Topics > "}
                {topic_name}
                {" > "}
            </Typography>
            <Card className={classes.postCard}>
                <CardContent>
                    <Box display="flex" alignItems="center" justifyContent="space-between" gap={2}>
                        <Typography variant="h5" className={classes.postTitle}>
                            {post.title}
                        </Typography>
                        <Box display="flex" gap={0.5} sx={{ alignSelf: "flex-start", mt: -0.5, mr: -1 }}>
                            <IconButton size="small" aria-label="edit topic" sx={{ p: 0.25 }}>
                                <ModeEditIcon fontSize="small" />
                            </IconButton>
                            <IconButton size="small" aria-label="delete topic" sx={{ p: 0.25 }}>
                                <DeleteIcon fontSize="small" />
                            </IconButton>
                        </Box>
                    </Box>
                    <Typography variant="body2" color="textPrimary" className={classes.postBody} component="p">
                        {post.content}
                    </Typography>
                </CardContent>
                <CardActions className={classes.cardActions}>
                    <Typography color="textSecondary" className={classes.metadata}>
                        {"Topic > " + post.topic_id}
                    </Typography>
                    <Typography color="textSecondary" className={classes.metadata}>
                        {"Posted by "}
                        <strong>{post.author}</strong>
                        {" on " + post.timestamp.toLocaleString()}
                    </Typography>
                </CardActions>
            </Card>
        </Box>
    );
};

export default PostItem;

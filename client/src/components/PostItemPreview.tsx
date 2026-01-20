import Post from "../types/Post";
import { makeStyles } from "@mui/styles";
import { Button, Card, CardActions, CardContent, Typography } from "@mui/material";
import React from "react";
import { Link } from "react-router-dom";

type Props = {
    post: Post;
    topicID: string;
};

const useStyles = makeStyles(() => ({
    postTitle: {
        fontSize: 24,
        whiteSpace: "pre-wrap",
        paddingBottom: "0.75em",
        textAlign: "left",
        fontWeight: 500,
    },
    postBody: {
        fontSize: 16,
        textAlign: "left",
        overflow: "hidden",
        textOverflow: "ellipsis",
        display: "-webkit-box",
        WebkitLineClamp: 2,
        WebkitBoxOrient: "vertical",
        lineHeight: 1.6,
        letterSpacing: "0.01em",
        wordBreak: "break-word",
    },
    postCard: {
        minWidth: 600,
        marginBottom: "0.75rem",
        padding: "6px",
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

const PostItemPreview: React.FC<Props> = ({ post, topicID }) => {
    const classes = useStyles();

    return (
        <>
            <Card className={classes.postCard}>
                <CardContent>
                    <Typography variant="h5" className={classes.postTitle}>
                        {post.title}
                    </Typography>
                    <Typography variant="body2" color="textPrimary" className={classes.postBody} component="p">
                        {post.content}
                    </Typography>
                </CardContent>
                <CardActions className={classes.cardActions}>
                    <Button size="small" component={Link} to={`/topics/${topicID}/posts/${post.id}`}>
                        View
                    </Button>
                    <Typography color="textSecondary" className={classes.metadata}>
                        {"Posted by "}
                        <strong>{post.author}</strong>
                        {" on " + post.timestamp.toLocaleString()}
                    </Typography>
                </CardActions>
            </Card>
        </>
    );
};

export default PostItemPreview;

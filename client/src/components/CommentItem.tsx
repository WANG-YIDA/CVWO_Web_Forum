import Comment from "../types/Comment";
import DeleteIcon from "@mui/icons-material/Delete";
import { Box, Card, CardActions, CardContent, IconButton, Typography } from "@mui/material";
import { makeStyles } from "@mui/styles";
import React from "react";

type Props = {
    comment: Comment;
};
const useStyles = makeStyles(() => ({
    commentBody: {
        fontSize: 16,
        whiteSpace: "pre-wrap",
        paddingBottom: "0.4em",
        textAlign: "left",
    },
    commentCard: {
        width: 800,
        marginBottom: "0.75rem",
        padding: "6px",
    },
    cardContent: {
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-start",
        gap: "0.5rem",
    },
    metadata: {
        fontSize: "12px !important",
        alignSelf: "flex-end",
        marginTop: "auto",
    },
    cardActions: {
        display: "flex",
        justifyContent: "flex-end",
        alignItems: "center",
        padding: "0 16px 12px",
    },
}));

const CommentItem: React.FC<Props> = ({ comment }) => {
    const classes = useStyles();

    return (
        <Card className={classes.commentCard}>
            <CardContent className={classes.cardContent}>
                <Box sx={{ width: "100%", maxWidth: 800, padding: "0 16px", margin: "0 auto", position: "relative" }}>
                    <IconButton
                        size="small"
                        aria-label="delete comment"
                        sx={{ p: 0.25, position: "absolute", top: -5, right: 25 }}
                    >
                        <DeleteIcon fontSize="small" />
                    </IconButton>
                    <Typography
                        variant="body2"
                        color="textPrimary"
                        sx={{ mt: 1, pr: 6 }}
                        className={classes.commentBody}
                        component="p"
                    >
                        {comment.body}
                    </Typography>
                </Box>
            </CardContent>
            <CardActions className={classes.cardActions}>
                <Typography color="textSecondary" className={classes.metadata} gutterBottom sx={{ mt: -0.25 }}>
                    {"Commented by "}
                    <strong>{comment.author}</strong>
                    {" on " + comment.timestamp.toLocaleString()}
                </Typography>
            </CardActions>
        </Card>
    );
};

export default CommentItem;

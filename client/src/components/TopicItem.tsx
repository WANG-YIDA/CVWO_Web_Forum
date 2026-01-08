import Topic from "../types/Topic";
import { makeStyles } from "@mui/styles";
import { Box, Button, Card, CardActions, CardContent, IconButton, Typography } from "@mui/material";
import ModeEditIcon from "@mui/icons-material/ModeEdit";
import DeleteIcon from "@mui/icons-material/Delete";
import React from "react";

type Props = {
    topic: Topic;
};

const useStyles = makeStyles(() => ({
    topicName: {
        fontSize: 32,
        whiteSpace: "pre-wrap",
        paddingBottom: "1em",
        textAlign: "left",
    },
    topicDescription: {
        fontSize: 16,
        whiteSpace: "pre-wrap",
        paddingBottom: "1em",
        textAlign: "left",
    },
    topicCard: {
        width: "100%",
        maxWidth: 800,
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

const TopicItem: React.FC<Props> = ({ topic }) => {
    const classes = useStyles();

    return (
        <Card className={classes.topicCard}>
            <CardContent>
                <Box display="flex" alignItems="center" justifyContent="space-between" gap={2}>
                    <Typography variant="h5" sx={{ pl: 1 }} className={classes.topicName}>
                        {topic.name}
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
                <Typography
                    variant="body2"
                    color="textPrimary"
                    sx={{ px: 3 }}
                    className={classes.topicDescription}
                    component="p"
                >
                    {topic.description}
                </Typography>
            </CardContent>
            <CardActions className={classes.cardActions}>
                <Button size="small" sx={{ pl: 1 }}>
                    View Posts
                </Button>
                <Typography color="textSecondary" sx={{ pr: 1 }} className={classes.metadata}>
                    {"Created by "}
                    <strong>{topic.author}</strong>
                    {" on " + topic.timestamp.toLocaleString()}
                </Typography>
            </CardActions>
        </Card>
    );
};

export default TopicItem;

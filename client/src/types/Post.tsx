type Post = {
    id: number;
    topic_id: number;
    title: string;
    content: string;
    user_id: number;
    author: string;
    timestamp: Date;
};

export default Post;

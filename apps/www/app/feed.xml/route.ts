import { authors } from "@/content/blog/authors";
import { type Post, allPosts } from "content-collections";
import RSS from "rss";
const feed = new RSS({
  title: "Ghost",
  description: "Open Source API Development platform",
  site_url: "https://ghost.com",
  feed_url: "https://ghost.com/feed.xml",
  copyright: `${new Date().getFullYear()} Ghost`,
  language: "en",
  pubDate: new Date(),
});

const posts = allPosts.sort((a: Post, b: Post) => {
  return new Date(b.date).getTime() - new Date(a.date).getTime();
});

export async function GET() {
  posts.map((post) => {
    const author = authors[post.author];
    feed.item({
      title: post.title,
      guid: `https://ghost.com/blog/${post.slug}`,
      url: `https://ghost.com/blog/${post.slug}`,
      date: post.date,
      description: post.description,
      author: author.name,
      categories: post.tags || [],
    });
  });

  return new Response(feed.xml({ indent: true }), {
    headers: {
      "Content-Type": "application/atom+xml; charset=utf-8",
    },
  });
}

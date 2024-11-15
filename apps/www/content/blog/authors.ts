export type Author = {
  name: string;
  role: string;
  image: {
    src: string;
    alt?: string;
  };
};

type Authors = {
  [key: string]: Author;
};

export const authors: Authors = {
  dom: {
    name: "Mohamed BenChaliah",
    role: "Engineer",
    image: { src: "/images/team/dom.jpeg" },
  },
  andreas: {
    name: "Developers",
    role: "Co-Founder",
    image: { src: "/images/team/andreas.jpeg" },
  },
  wilfred: {
    name: "Mohamed BenChaliah",
    role: "Freelance Writer",
    image: { src: "/images/blog-images/ocr-post/wilfred.jpg" },
  },
  michael: {
    name: "Mohamed BenChaliah",
    role: "Developer",
    image: { src: "/images/team/michael.jpg" },
  },
};

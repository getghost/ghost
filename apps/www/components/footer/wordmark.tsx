"use client";

import { motion, useReducedMotion } from "framer-motion";
import type React from "react";

function FadeInStagger({ ...props }) {
  return (
    <motion.div
      initial="hidden"
      whileInView="visible"
      viewport={{ once: false, margin: "0px 0px 0px 0px" }}
      transition={{ staggerChildren: 0.15 }}
      {...props}
    />
  );
}
export const Wordmark: React.FC<{ className?: string }> = ({ className }) => {
  const shouldReduceMotion = useReducedMotion();

  const variants = {
    hidden: { opacity: 0, y: shouldReduceMotion ? 0 : 64 },
    visible: { opacity: 1, y: 0 },
  };
  const transition = {
    duration: 0.05,
    ease: "easeOut",
    type: "spring",
    stiffness: 200,
    damping: 50,
  };

  return (
    <FadeInStagger className={className}>
      <svg
        width="1376"
        height="248"
        viewBox="0 0 1376 248"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
      </svg>
    </FadeInStagger>
  );
};

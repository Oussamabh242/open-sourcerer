let styleMap = {
  H1: "text-3xl font-bold text-gray-950 mb-4 dark:text-gray-50", // Darkest gray -> light gray
  H2: "text-2xl font-semibold text-gray-900 mb-3 dark:text-gray-100", // Slightly lighter gray -> light gray
  H3: "text-xl font-medium text-gray-800 mb-2 dark:text-gray-200", // Medium gray -> lighter gray
  H4: "text-lg font-medium text-gray-600 mb-2 dark:text-gray-300", // Lighter gray -> lighter gray in dark mode
  H5: "text-base font-medium text-gray-500 mb-1 dark:text-gray-400", // Even lighter gray -> similar contrast
  H6: "text-sm font-medium text-gray-400 mb-1 dark:text-gray-500", // Adjusted for better visibility
  P: "mb-4 leading-relaxed dark:text-gray-100 text-base", // Medium gray -> lighter text for readability
  STRONG: "font-bold text-gray-900 dark:text-white", // Darkest gray -> white in dark mode
  EM: "italic text-gray-600 dark:text-gray-400", // Adjust for dark mode
  A: "text-blue-700 hover:text-blue-900 underline dark:text-blue-400 dark:hover:text-blue-600", // Lighter blues in dark mode
  CODE: "bg-gray-200 text-gray-800 px-1 py-0.5 rounded dark:bg-gray-800 dark:text-gray-200", // Swap to dark background with light text
  PRE_CODE: "bg-gray-200 text-gray-800 p-4 rounded-lg mb-4 overflow-x-auto dark:bg-gray-900 dark:text-gray-100", // Darker background for code blocks
  UL_LI: "list-disc list-inside my-2 text-white dark:text-gray-400", // Adjust list items for dark mode
  OL_LI: "list-decimal list-inside my-2 text-gray-600 dark:text-gray-400", // Adjust for ordered lists in dark mode
  BLOCKQUOTE: "border-l-4 border-gray-400 pl-4 italic text-gray-500 mb-4 dark:border-gray-600 dark:text-gray-300", // Border and text color for dark mode
  HR: "border-t border-gray-400 my-8 dark:border-gray-600", // Adjust horizontal rule for dark mode
  IMG: "max-w-full h-auto rounded-lg shadow-md dark:shadow-none", // Optional: Remove shadow in dark mode
  TABLE: "border-collapse border border-gray-400 mb-4 dark:border-gray-600", // Adjust table border for dark mode
  TH: "bg-gray-100 font-semibold p-2 border border-gray-300 text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-300", // Adjust table header for dark mode
  TD: "p-2 border border-gray-300 text-gray-700 dark:border-gray-700 dark:text-gray-300", // Adjust table cells for dark mode
  TASK_LIST: "list-none pl-6 relative before:content-['✓'] before:absolute before:left-0 before:text-gray-500 dark:before:text-gray-400", // Task list styling for dark mode
  MARK: "bg-yellow-200 px-1 rounded dark:bg-yellow-700", // Adjust highlight color for dark mode
  SUP: "text-xs align-super text-gray-500 dark:text-gray-400", // Adjust superscript for dark mode
  DEL: "line-through text-gray-400 dark:text-gray-600", // Strike-through adjustment for dark mode
  DL_DT: "font-bold mb-1 text-gray-800 dark:text-gray-200", // Definition term adjustment
  DL_DD: "pl-4 mb-4 text-gray-600 dark:text-gray-400", // Definition description adjustment
  ABBR: "underline dotted text-gray-600 dark:text-gray-400", // Abbreviation adjustment
  SUB: "text-xs align-sub text-gray-600 dark:text-gray-400" // Subscript adjustment
};



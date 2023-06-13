import { FC, useRef } from "react";

import AutoSizer from "react-virtualized-auto-sizer";
import type { EditorProps } from "@monaco-editor/react";
import MonacoEditor from "@monaco-editor/react";

// for reference:
// https://github.com/Microsoft/vscode/blob/913e891c34f8b4fe2c0767ec9f8bfd3b9dbe30d9/src/vs/editor/standalone/common/themes.ts#L13
const beforeMount: EditorProps["beforeMount"] = (monaco) => {
  monaco.editor.defineTheme("direktiv-dark", {
    base: "vs-dark",
    inherit: true,
    rules: [
      {
        foreground: "505050", // grayDark.gray8
        fontStyle: "italic",
        token: "comment",
      },
      {
        foreground: "6473FF", // primary.400
        token: "number",
      },
      {
        foreground: "6473FF", // primary.400
        token: "string.yaml",
      },
      {
        foreground: "7E7E7E", // grayDark.gray10
        token: "type",
      },
    ],
    colors: {
      "editor.background": "#000000",
      "editor.selectionBackground": "#ffffff2e", // whiteA.whiteA7
    },
  });

  monaco.editor.defineTheme("direktiv-light", {
    base: "vs",
    inherit: true,
    rules: [
      {
        foreground: "C7C7C7", // gray.gray8
        fontStyle: "italic",
        token: "comment",
      },
      {
        foreground: "5364FF", // primary.500
        token: "number",
      },
      {
        foreground: "5364FF", // primary.500
        token: "string.yaml",
      },
      {
        foreground: "858585", // gray.gray10
        token: "type",
      },
      {
        foreground: "858585",
        token: "string.key.json", //JSON Key
      },
      {
        foreground: "5364FF",
        token: "string.value.json", //JSON Value
      },
      {
        foreground: "858585",
        token: "tag", //HTML Tag
      },
      {
        foreground: "858585",
        token: "metatag.html", //HTML Meta tag
      },
      {
        foreground: "5364FF",
        token: "metatag.content.html", //HTML Meta tag content
      },
      {
        foreground: "5364FF",
        token: "delimiter", //HTML Meta tag content
      },
      {
        foreground: "858585",
        token: "attribute.name", //HTML Attribute Name
      },
      {
        foreground: "5364FF",
        token: "attribute.value.html", //HTML Attribute Name
      },
    ],
    colors: {
      "editor.foreground": "#858585", // added for HTML Tag Content
      "editor.background": "#ffffff",
      "editor.selectionBackground": "#00000012", // blackA.blackA4
    },
  });
};

type EditorType = Parameters<NonNullable<EditorProps["onMount"]>>[0];

const Editor: FC<
  Omit<EditorProps, "beforeMount" | "onMount" | "onChange"> & {
    theme?: "light" | "dark";
    onSave?: (value: string | undefined) => void;
    onChange?: (value: string | undefined) => void;
    language?: "html" | "css" | "json" | "shell" | "plaintext" | "yaml";
  }
> = ({ options, theme, onSave, onChange, language = "yaml", ...props }) => {
  const monacoRef = useRef<EditorType>();

  const handleChange = () => {
    onChange && onChange(monacoRef.current?.getValue());
  };

  const onMount: EditorProps["onMount"] = (editor, monaco) => {
    monacoRef.current = editor;
    editor.focus();
    monacoRef.current.onDidChangeModelContent(handleChange);

    onSave &&
      editor.addCommand(monaco.KeyMod.CtrlCmd | monaco.KeyCode.KeyS, () => {
        onSave(
          monacoRef.current?.getValue()
            ? `${monacoRef.current?.getValue()}`
            : undefined
        );
      });
  };

  return (
    <AutoSizer>
      {({ height, width }) => (
        <MonacoEditor
          width={width}
          height={height}
          beforeMount={beforeMount}
          onMount={onMount}
          options={{
            scrollBeyondLastLine: false,
            cursorBlinking: "smooth",
            wordWrap: true,
            fontSize: "13px",
            minimap: {
              enabled: false,
            },
            contextmenu: false,
            ...options,
          }}
          loading=""
          language={language}
          theme={theme === "dark" ? "direktiv-dark" : "direktiv-light"}
          {...props}
        />
      )}
    </AutoSizer>
  );
};

export default Editor;

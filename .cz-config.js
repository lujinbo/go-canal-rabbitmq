/*
 *  ┌─────────────────────────────────────────────────────────────┐
 *  │┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┐│
 *  ││ Esc  │  !1  │  @2  │  #3  │  $4  │  %5  │  ^6  │  &7  │  *8  │  (9  │  )0  │  _-  │  +=  │  |\  │  `~  ││
 *  │├───┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴───┤│
 *  ││ Tab      │  Q   │  W   │  E   │  R   │  T   │  Y   │  U   │  I   │  O   │   P  │  {[  │  }]  │      BS  ││
 *  │├─────┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴┬──┴─────┤│
 *  ││ Ctrl       │   A  │   S  │   D  │   F  │   G  │   H  │   J  │   K  │   L  │  : ; │  " ' │         Enter  ││
 *  │├──────┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴─┬─┴────┬───┤│
 *  ││ Shift          │   Z  │   X  │   C  │   V  │   B  │   N  │   M  │  < , │  > . │  ? / │      Shift │   Fn ││
 *  │└─────┬──┴┬──┴──┬┴───┴───┴───┴───┴───┴──┬┴───┴┬──┴┬─────┴───┘│
 *  │            │  Fn  │   Alt    │                     Space                    │    Alt   │  Win │            Goldbug │
 *  │            └───┴─────┴───────────────────────┴─────┴───┘                    │
 *  └─────────────────────────────────────────────────────────────┘
 * @Author: 刘白菜 liushuai.baicai@hotmail.com
 * @Date: 2022-07-18 15:34:43
 * @LastEditors: 刘白菜
 * @LastEditTime: 2022-08-05 12:47:10
 * @Description: commitizen提交规范定义文件
 */

module.exports = {
  // type 类型（定义之后，可通过上下键选择）
  types: [
    { value: "✨", name: "feat:     新增功能" },
    { value: "🐛", name: "fix:      修复 bug" },
    { value: "🎨", name: "style:    UI样式调整" },
    { value: "⚡️", name: "perf:     性能优化" },
    { value: "🚧", name: "format:   代码格式（不影响功能，例如空格、分号等格式修正）" },
    { value: "📝", name: "docs:     文档变更" },
    { value: "🔧", name: "tools:    构建工具、脚本、项目配置等" },
    { value: "➖", name: "revert:   回滚 commit" },
  ],

  // scope 类型（定义之后，可通过上下键选择）
  scopes: [
    ["平台", "蜜罐相关"],
    ["大屏", "大屏相关"],
    ["配置", "项目配置相关"],
    ["其他", "其他修改"],
    // 如果选择 custom，后面会让你再输入一个自定义的 scope。也可以不设置此项，把后面的 allowCustomScopes 设置为 true
    ["custom", "以上都不是？我要自定义"],
  ].map(([value, description]) => {
    return {
      value,
      name: `${value.padEnd(30)} (${description})`,
    };
  }),

  // 是否允许自定义填写 scope，在 scope 选择的时候，会有 empty 和 custom 可以选择。
  // allowCustomScopes: true,

  // allowTicketNumber: false,
  // isTicketNumberRequired: false,
  // ticketNumberPrefix: 'TICKET-',
  // ticketNumberRegExp: '\\d{1,5}',

  // 针对每一个 type 去定义对应的 scopes，例如 fix
  /*
    scopeOverrides: {
      fix: [
        { name: 'merge' },
        { name: 'style' },
        { name: 'e2eTest' },
        { name: 'unitTest' }
      ]
    },
    */

  // 交互提示信息
  messages: {
    type: "确保本次提交遵循规范！\n选择你要提交的类型：",
    scope: "\n选择一个 scope（可选）：",
    // 选择 scope: custom 时会出下面的提示
    customScope: "请输入自定义的 scope：",
    subject: "填写简短精炼的变更描述：\n",
    body: '填写更加详细的变更描述（可选）。使用 "|" 换行：\n',
    breaking: "列举非兼容性重大的变更（可选）：\n",
    footer: "列举出所有变更的 JIRA Resolved（可选）。 例如: MG-31, MG-34：\n",
    confirmCommit: "确认提交？",
  },

  // 设置只有 type 选择了 feat 或 fix，才询问 breaking message
  allowBreakingChanges: ["feat", "fix"],

  // 跳过要询问的步骤
  skipQuestions: ["body"],

  // subject 限制长度
  subjectLimit: 100,
  breaklineChar: "|", // 支持 body 和 footer
  footerPrefix: "JIRA Resolved:",
  // askForBreakingChangeFirst : true,
};

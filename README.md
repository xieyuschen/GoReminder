# GoReminder
关于本项目，本项目是使用Go实现的小说更新提示器。这次的项目关注点主要有以下几点：  
- 扩充可订阅的小说范围  
过去只能订阅一本小说，现在期望可以进行一个小说的订阅管理。
- 充分使用Go的Concurrency功能
选择Go的原因就是因为Go提供给了用户强大易用的并发功能，而该项目基于一个网络爬虫与邮件发送，与网络服务密不可分。是一个很好的使用并发的例子。但是需要我仔细的设计并发。并且可能在未来的学习中进行反复的版本迭代写出效率更高更加稳固的代码:)
- *具有生成整本小说的功能  
属于随后的扩展功能，用于拓展适用范围。  

# 整体设计
代码结构分为以下若干块：  
- Web爬虫  
完成对Html的解析，获取最新的章节，获取最新章节的内容与标题。 
- 邮件系统  
给订阅的用户发送一封邮件，包含更新小说的标题与文章内容。
- 数据库设计  
内容的存储直接放在数据库中，而不是放在一个文件当中。  
- 处理更新与邮件发送  
实时扫描，在扫描的时候能够灵活运用**Concurrency**的相关知识实现一个高效的管理系统。  

# 具体设计
## Web爬虫：
Web爬虫即正常的使用`Html parser`解析对应的html界面，然后获取到对应内容即可。  
- 多本小说  
多本小说即需要保存对应的`url`链接作为map的key，然后扫描的时候获取的`章节:内容`信息存入到map中去，然后供邮件进行使用。

## Reminder组件
Reminder组件的设计较为复杂，这里简单的写一下。需要完成的功能模块：  
- 持续扫描小说是否更新  
扫描毫无疑问应该在`main goroutine`上进行，他完成对整个程序逻辑的组装，将不同组件组装在一起。
```
Start
|
+-->check whether the subsribed novel is updated
^   {+-->Send a get request and get novelInfo
|    +-->Query last-sending email in database
|    +-->Whether db chapter is newest?+--yes-->Continue this scan loop   
|                                     |
|                 |Block until       no happened
Next loop         |                   |  
|                 +-->Get new chapter's content   
|                 |
|                 +-->Send Email to Subsriber,update db
|                 |
+-----------------|   
```
需要考虑的是Go routine的阻塞问题，直接使用一个`Unbuffered channel`即可。






















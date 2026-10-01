---
tags:
  - Operating_System
  - Linux
  - linux-torvalds
  - open-source
date: 10-30-2025
---
# [Linux](https://en.wikipedia.org/wiki/Linux)

![Linux|150](https://upload.wikimedia.org/wikipedia/commons/3/35/Tux.svg)

*Imported from [Wikipedia](https://en.wikipedia.org/)*

> [!info]- Quick facts
> [[Tux (mascot)|Tux]], the mascot of Linux, created by [[Larry Ewing]][^1]
> **[[Programmer|Developer]]:** [[Linus Torvalds]] ([[Benevolent dictator for life|BDFL]]); Community contributors
> **OS family:** [[Unix-like]]
> **Working state:** Current
> **Source model:** [[Open source]] ([[Linux Kernel|kernel]])
> **Initial release:** 17 September 1991 (1991-09-17) (kernel)[^2]
> **Marketing target:** General-purpose
> **Available in:** Multilingual
> **Supported platforms:** Various
> **[[Kernel (operating system)|Kernel]] type:** [[Monolithic kernel|Monolithic]]
> **[[User space and kernel space|Userland]]:** Typically [[Util-linux]] and one of [[GNU Core Utilities|coreutils]], [[BusyBox]], or [[Toybox]]
> **Influenced by:** [[Unix]] and [[Minix]]
> **[[Software license|License]]:** [[GPLv2]] ([[Linux Kernel|kernel]])[^9]
> **Articles in the series**
> [[Linux Kernel]]; [[Linux distribution]]

**Linux** ( )[^11] is a family of [[Free and open-source software]] [[Unix-like]] operating systems based on the [[Linux Kernel]],[^12] which was first released on 17 September 1991 by [[Linus Torvalds]].[^13][^14] Some members of the family are typically [[Package manager|packaged]] as a [[Linux distribution|distribution]] (a.k.a. *distro*), which includes the [[Kernel (operating system)|kernel]] alongside supporting system software and [[Library (computing)|libraries]] developed by third parties—such as [[GNU Project|GNU]], [[Red Hat]], and [[X.Org Foundation|X.Org]]—to create a complete operating system; however, not all Linux-based operating systems are considered distros, with [[Android (operating system)|Android]] being an example. Linux was originally designed as a clone of [[Unix]] and is distributed under the copyleft [[GPL]] license.[^15]

There are many [[List of Linux distributions|thousands of Linux distributions]], many based directly or indirectly on other distributions;[^16][^17] popular Linux distros[^16][^18][^19] include [[Debian]], [[Fedora Linux]], [[Linux Mint]], [[Arch Linux]], and [[Ubuntu]], while commercial distributions include [[Red Hat Enterprise Linux]], [[SUSE Linux Enterprise]], and [[ChromeOS]]. Linux distributions are frequently used in server platforms.[^20][^21] Many Linux distributions use the word "**Linux**" in their name, but the [[Free Software Foundation]] uses and recommends the name "**GNU/Linux**", to emphasize the use and importance of [[GNU]] software in many distributions. There is [[GNU/Linux naming controversy|controversy surrounding this]].[^22][^23] Other than the [[Linux Kernel]], key components that make up a distribution may include a [[Display server|display server (windowing system)]], a [[Package manager]], a [[Bootloader]], and a [[Unix shell]].

Linux is one of the most prominent examples of [[Free and open-source software|free and open-source]] software collaboration. The Linux kernel is considered by many to be the largest free and open source project. While originally developed for [[X86]]-based [[Personal computer|personal computers]], it has since been [[Porting|ported]] to more [[Computer hardware platforms|platforms]] than any other operating system.[^24] When combined with [[Android (operating system)|Android]], which uses a Linux-based kernel and is designed for [[Smartphone|smartphones]], they have the [[Usage share of operating systems|largest installed base]] of all [[General-purpose operating system|general-purpose operating systems]].

## Overview

The [[Linux Kernel]] was created by [[Linus Torvalds]], following the lack of a working [[Kernel (operating system)|kernel]] for [[GNU]], a [[Unix]]-compatible operating system made entirely of [[Free software]] that had been in development since 1983 by the [[GNU Project]], led by [[Richard Stallman]]. A working Unix system called [[Minix]] was later released but its license was not entirely free at the time[^25] and it was made for education purposes. The first entirely free Unix for personal computers, [[386BSD]], did not appear until 1992, by which time Torvalds had already built and publicly released the first version of the Linux kernel on the [[Internet]].[^26] Like GNU and 386BSD, Linux did not have any Unix code, being a fresh re-implementation, and therefore avoided [[UNIX System Laboratories, Inc. v. Berkeley Software Design, Inc.|legal issues]] from [[AT&T Corporation|AT&T]].[^27] Linux distributions became popular in the 1990s and made Unix technologies accessible to home users on personal computers whereas previously it had been confined to sophisticated [[Workstations]].[^28]

Desktop Linux distributions include a [[Windowing system]] such as [[X Window System|X11]] or [[Wayland (protocol)|Wayland]] and a [[Desktop environment]] such as [[KDE Plasma]], [[GNOME]], or [[Xfce]]. Distributions intended for [[Server (computing)|servers]] may not have a [[Graphical user interface]] at all or include a [[Solution stack]] such as [[LAMP (software bundle)|LAMP]].

The [[Source code]] of Linux may be used, modified, and distributed commercially or non-commercially by anyone under the terms of its respective licenses, such as the [[GNU General Public License]] (GPL). The license means creating novel distributions is permitted by anyone[^29] and is easier than it would be for an operating system such as [[macOS]] or [[Microsoft Windows]].[^30][^31][^32] The Linux kernel, for example, is licensed under the GPLv2, with an exception for [[System call|system calls]] that allows code that calls the kernel via system calls not to be licensed under the GPL.[^33][^29]

Because of the dominance of Linux-based [[Android (operating system)|Android]] on [[Smartphone|smartphones]], Linux, including Android, has the [[Usage share of operating systems|largest installed base]] of all [[General-purpose operating system|general-purpose operating systems]] as of May 2022.[^34][^35][^36]
Linux is, as of March 2024, used by around 4 percent of [[Desktop computer|desktop computers]].[^37] The [[Chromebook]], which runs the Linux kernel-based [[ChromeOS]],[^38][^39] dominates the US [[K–12]] education market and represents nearly 20 percent of sub-$300 [[Laptop]] sales in the US.[^40] Linux is the leading operating system on servers (over 96.4% of the top one million web servers' operating systems are Linux),[^41] leads other [[Big iron (computing)|big iron]] systems such as [[Mainframe computer|mainframe computers]],[^42] and is used on all of the [[TOP500|world's 500 fastest supercomputers]] (as of November 2017, having gradually displaced all competitors).[^43][^44]

Linux also runs on [[Embedded system|embedded systems]], i.e., devices whose operating system is typically built into the [[Firmware]] and is highly tailored to the system. This includes [[Router (computing)|routers]], [[Automation]] controls, [[Smart home technology|smart home devices]], [[Video game console|video game consoles]], [[Television|televisions]] (Samsung and LG [[Smart TV|smart TVs]]),[^45][^46][^47] [[Automobiles]] (Tesla, Audi, Mercedes-Benz, Hyundai, and Toyota),[^48] and [[Spacecraft]] ([[Falcon 9]] rocket, [[SpaceX Dragon 2|Dragon]] crew capsule, and the [[Ingenuity (helicopter)|Ingenuity]] Mars helicopter).[^49][^50]

## History

### Precursors

![Linus Torvalds|190](https://upload.wikimedia.org/wikipedia/commons/5/5c/Linus_Torvalds_%28cropped%29.jpg?utm_source=en.wikipedia.org&utm_campaign=parser&utm_content=thumbnail_unscaled)
*[[Linus Torvalds]], principal author of the Linux kernel*

The [[Unix]] operating system was conceived of and implemented in 1969, at [[AT&T Corporation|AT&T]]'s [[Bell Labs]] in the United States, by [[Ken Thompson]], [[Dennis Ritchie]], [[Douglas McIlroy]], and [[Joe Ossanna]].[^51] First released in 1971, Unix was written entirely in [[Assembly language]], as was common practice at the time. In 1973, in a key pioneering approach, it was rewritten in the [[C (programming language)|C]] programming language by Dennis Ritchie (except for some hardware and I/O routines). The availability of a [[High-level language]] implementation of Unix made its [[Porting]] to different computer platforms easier.[^52]

As a 1956 [[Antitrust case]] forbade [[AT&T]] from entering the computer business,[^53] AT&T provided the operating system's [[Source code]] to anyone who asked. As a result, Unix use grew quickly and it became widely adopted by [[Academic institution|academic institutions]] and businesses. In 1984, [[Breakup of the Bell System|AT&T divested itself]] of its [[Regional Bell Operating Company|regional operating companies]], and was released from its obligation not to enter the computer business; freed of that obligation, Bell Labs began selling Unix as a [[Proprietary software|proprietary]] product, where users were not legally allowed to modify it.[^54][^55]

[[Onyx Systems]] began selling early microcomputer-based Unix workstations in 1980. Later, [[Sun Microsystems]], founded as a spin-off of a student project at [[Stanford University]], also began selling Unix-based desktop workstations in 1982. While Sun workstations did not use commodity PC hardware, for which Linux was later originally developed, it represented the first successful commercial attempt at distributing a primarily single-user microcomputer that ran a Unix operating system.[^56][^57]

With Unix increasingly "locked in" as a proprietary product, the [[GNU Project]], started in 1983 by [[Richard Stallman]], had the goal of creating a "complete Unix-compatible [[Software system]]" composed entirely of [[Free software]]. Work began in 1984.[^58] Later, in 1985, Stallman started the [[Free Software Foundation]] and wrote the [[GNU General Public License]] (GNU GPL) in 1989. By the early 1990s, many of the programs required in an operating system (such as libraries, [[Compiler|compilers]], [[Text editor|text editors]], a [[Shell (computing)|command-line shell]], and a [[Windowing system]]) were completed, although low-level elements such as [[Device driver|device drivers]], [[Daemon (computing)|daemons]], and the [[Kernel (operating system)|kernel]], called [[GNU Hurd]], were stalled and incomplete.[^59]

[[Minix]] was created by [[Andrew S. Tanenbaum]], a [[Computer science]] professor, and released in 1987 as a minimal Unix-like operating system targeted at students and others who wanted to learn operating system principles. Although the [[Source-available software|complete source code of Minix was freely available]], the licensing terms prevented it from being [[Free software]] until the licensing changed in April 2000.[^60]

### Creation

While attending the [[University of Helsinki]] in the fall of 1990, Torvalds enrolled in a Unix course.[^61] The course used a [[MicroVAX]] [[Minicomputer]] running [[Ultrix]], and one of the required texts was *[[Operating Systems: Design and Implementation]]* by [[Andrew S. Tanenbaum]]. This textbook included a copy of Tanenbaum's [[Minix]] operating system. It was with this course that Torvalds first became exposed to Unix. In 1991, he became curious about operating systems.[^62] Frustrated by the licensing of Minix, which at the time limited it to educational use only,[^60] he began to work on his operating system kernel, which eventually became the Linux kernel.

On 3 July 1991, to implement Unix [[System call|system calls]], [[Linus Torvalds]] attempted unsuccessfully to obtain a [[Digital copy]] of the [[POSIX]] standards [[Software documentation|documentation]] with a request to the *comp.os.minix* [[Usenet newsgroup|newsgroup]].[^63] After not finding the POSIX documentation, Torvalds initially resorted to determining system calls from [[SunOS]] documentation owned by the university for use in operating its [[Sun Microsystems]] server. He also learned some system calls from Tanenbaum's Minix text.

Torvalds began the development of the [[Linux Kernel]] on [[Minix]] and applications written for Minix were also used on Linux. Later, Linux matured and further Linux [[Kernel (operating system)|kernel]] development took place on Linux systems.[^64] GNU applications also replaced all Minix components, because it was advantageous to use the freely available code from the GNU Project with the fledgling operating system; code licensed under the GNU GPL can be reused in other [[Computer program|computer programs]] as long as they also are released under the same or a compatible license. Torvalds initiated a switch from his original license, which prohibited commercial redistribution, to the GNU GPL.[^65] Developers worked to integrate GNU components with the Linux kernel, creating a fully functional and free [[Operating system]].[^66]

Although not released until 1992, due to [[UNIX System Laboratories, Inc. v. Berkeley Software Design, Inc.|legal complications]], the development of [[386BSD]], from which [[NetBSD]], [[OpenBSD]] and [[FreeBSD]] descended, predated that of Linux. Linus Torvalds has stated that if the [[GNU kernel]] or 386BSD had been available in 1991, he probably would not have created Linux.[^67][^25]

### Copyright, trademark, and naming

![image|250](https://upload.wikimedia.org/wikipedia/commons/c/c9/Gnulinux.svg)
*[[Tux (mascot)|Tux]] is sometimes stylized with incorporation of the [[GNU]] logo*

The Linux kernel is [[Software license|licensed]] under the GNU General Public License (GPL), version 2. The GPL requires that anyone who distributes software based on source code under this license must make the originating source code (and any modifications) available to the recipient under the same terms.[^68] Other key components of a typical Linux distribution are also mainly licensed under the GPL, but they may use other licenses; many libraries use the [[GNU Lesser General Public License]] (LGPL), a more permissive variant of the GPL, and the [[X.Org Server|X.Org]] implementation of the X Window System uses the [[MIT License]].

Torvalds states that the Linux kernel will not move from version 2 of the GPL to version 3.[^69][^70] He specifically dislikes some provisions in the new license which prohibit the use of the software in [[Digital rights management]].[^71] It would also be impractical to obtain permission from all the copyright holders, who number in the thousands.[^72]

A 2001 study of [[Red Hat Linux]] 7.1 found that this distribution contained 30 million [[Source lines of code|lines of source code]].[^73] Using the [[COCOMO|Constructive Cost Model]], the study estimated that this distribution required about eight thousand person-years of development time. According to the study, if all this software had been developed by conventional proprietary means, it would have cost about US$1.86 billion[^74] to develop in 2024 in the United States.[^73] Most of the source code (71%) was written in the C programming language, but many other languages were used, including [[C++]], [[Lisp (programming language)|Lisp]], assembly language, Perl, Python, [[Fortran]], and various [[Shell script|shell scripting]] languages. Slightly over half of all lines of code were licensed under the GPL. The Linux kernel itself was 2.4 million lines of code, or 8% of the total.[^73]

In a later study, the same analysis was performed for Debian version 4.0 (etch, which was released in 2007).[^75] This distribution contained close to 283 million lines of source code, and the study estimated that it would have required about seventy three thousand man-years and cost US$10.4 billion[^74] (in 2024 dollars) to develop by conventional means.

![image|250](https://upload.wikimedia.org/wikipedia/commons/thumb/e/ee/Linux_0_12.jpg/250px-Linux_0_12.jpg?utm_source=en.wikipedia.org&utm_campaign=parser&utm_content=thumbnail)
*5.25-inch [[Floppy disk|floppy disks]] holding a very early version of Linux*

In the United States, the name *Linux* is a trademark registered to Linus Torvalds.[^10] Initially, nobody registered it. However, on 15 August 1994, William R. Della Croce Jr. filed for the trademark *Linux*, and then demanded royalties from Linux distributors. In 1996, Torvalds and some affected organizations sued him to have the trademark assigned to Torvalds, and, in 1997, the case was settled.[^76] The licensing of the trademark has since been handled by the [[Linux Mark Institute]] (LMI). Torvalds has stated that he trademarked the name only to prevent someone else from using it. LMI originally charged a nominal sublicensing fee for use of the Linux name as part of trademarks,[^77] but later changed this in favor of offering a free, perpetual worldwide sublicense.[^78]

The Free Software Foundation (FSF) prefers *GNU/Linux* as the name when referring to the operating system as a whole, because it considers Linux distributions to be [[GNU variants|variants]] of the GNU operating system initiated in 1983 by [[Richard Stallman]], president of the FSF.[^22][^23] The foundation explicitly takes no issue over the name Android for the Android OS, which is also an operating system based on the Linux kernel, as GNU is not a part of it.

A minority of public figures and software projects other than Stallman and the FSF, notably distributions consisting of only free software, such as Debian (which had been sponsored by the FSF up to 1996),[^79] also use *GNU/Linux* when referring to the operating system as a whole.[^80][^81][^82] Most media and common usage, however, refers to this family of operating systems simply as *Linux*, as do many large Linux distributions (for example, [[SUSE Linux]] and [[Red Hat Enterprise Linux]]).

As of May 2011, about 8% to 13% of the [[Source lines of code|lines of code]] of the Linux distribution Ubuntu (version "Natty") is made of GNU components (the range depending on whether GNOME is considered part of GNU); meanwhile, 6% is taken by the Linux kernel, increased to 9% when including its direct dependencies.[^83]

#### Naming

Linus Torvalds had wanted to call his invention "Freax", a [[Portmanteau]] of "free", "freak", and "x" (as an allusion to Unix). During the start of his work on the system, some of the project's [[Makefile|makefiles]] included the name "Freax" for about half a year. Torvalds considered the name "Linux" but dismissed it as too egotistical.[^84]

To facilitate development, the files were uploaded to the [[File Transfer Protocol|FTP server]] of [[FUNET]] in September 1991. Ari Lemmke, Torvalds' coworker at the [[Helsinki University of Technology]] (HUT) who was one of the volunteer administrators for the FTP server at the time, did not think that "Freax" was a good name, so he named the project "Linux" on the server without consulting Torvalds.[^84] Later, however, Torvalds consented to "Linux".

According to a [[Usenet newsgroup|newsgroup]] post by Torvalds,[^11] the word "Linux" should be pronounced (  ) with a short 'i' as in 'print' and 'u' as in 'put'. To further demonstrate how the word "Linux" should be pronounced, he included an audio guide with the kernel source code.[^85] However, in this recording, he pronounces Linux as */ˈlinʊks/* () with a short but [[Close front unrounded vowel]], instead of a [[Near-close near-front unrounded vowel]] as in his newsgroup post.

### Commercial and popular uptake

The adoption of Linux in production environments, rather than being used only by hobbyists, started to take off first in the mid-1990s in the supercomputing community, where organizations such as [[NASA]] started replacing their increasingly expensive machines with [[Computer cluster|clusters]] of inexpensive commodity computers running Linux. Commercial use began when [[Dell]] and [[IBM]], followed by [[Hewlett-Packard]], started offering Linux support to escape [[Microsoft]]'s monopoly in the desktop operating system market.[^86]

Today, Linux systems are used throughout computing, from [[Embedded system|embedded systems]] to virtually all [[Supercomputer|supercomputers]],[^44][^87] and have secured a place in server installations such as the popular [[LAMP (software bundle)|LAMP]] application stack. The use of Linux distributions in home and enterprise desktops has been growing.[^88][^89][^90][^91][^92][^93][^94]

Linux distributions have also become popular in the [[Netbook]] market, with many devices shipping with customized Linux distributions installed, and Google releasing their own [[ChromeOS]] designed for netbooks.

Linux's greatest success in the consumer market is perhaps the mobile device market, with Android being the dominant operating system on [[Smartphone|smartphones]] and very popular on [[Tablet computer|tablets]] and, more recently, on [[Wearable technology|wearables]], and vehicles. [[Linux gaming]] is also on the rise with [[Valve Corporation|Valve]] showing its support for Linux and rolling out [[SteamOS]], its own gaming-oriented Linux distribution, which was later implemented in their [[Steam Deck]] platform. Linux distributions have also gained popularity with various local and national governments, such as the federal government of [[Brazil]].[^95]

### Development

Linus Torvalds is the lead maintainer for the Linux kernel and guides its development, while [[Greg Kroah-Hartman]] is the lead maintainer for the stable branch.[^96] Zoë Kooyman is the executive director of the Free Software Foundation,[^97] which in turn supports the GNU components.[^98] Finally, individuals and corporations develop third-party non-GNU components. These third-party components comprise a vast body of work and may include both kernel modules and user applications and libraries.

Linux vendors and communities combine and distribute the kernel, GNU components, and non-GNU components, with additional [[Package management]] software in the form of Linux distributions.

## Usage

### Market share and uptake

Many quantitative studies of free/open-source software focus on topics including market share and reliability, with numerous studies specifically examining Linux.[^99] The Linux market is growing, and the Linux operating system market size is expected to see a growth of 19.2% by 2027, reaching $15.64 billion, compared to $3.89 billion in 2019.[^100] Analysts project a Compound Annual Growth Rate (CAGR) of 13.7% between 2024 and 2032, culminating in a market size of US$34.90 billion by the latter year. Analysts and proponents attribute the relative success of Linux to its security, reliability, low cost, and freedom from [[Vendor lock-in]].[^101][^102]

Desktops and laptops
According to [[Web analytics|web server statistics]] (that is, based on the numbers recorded from visits to websites by client devices), in October 2024, the estimated market share of Linux on [[Desktop computer|desktop computers]] was around 4.3%. In comparison, [[Microsoft Windows]] had a market share of around 73.4%, while [[macOS]] covered around 15.5%.[^37]
Web servers
W3Cook publishes stats that use the top 1,000,000 Alexa domains,[^103] which as of May 2015 estimate that 96.55% of web servers run Linux, 1.73% run Windows, and 1.72% run FreeBSD.[^41]
W3Techs publishes stats that use the top 10,000,000 Alexa domains and the top 1,000,000 Tranco domains, updated monthly[^104] and as of November 2020 estimate that Linux is used by 39% of the web servers, versus 21.9% using [[Microsoft Windows]].[^105] 40.1% used other types of Unix.[^106]
[[International Data Corporation|IDC]]'s Q1 2007 report indicated that Linux held 12.7% of the overall server market at that time;[^107] this estimate was based on the number of Linux servers sold by various companies, and did not include server hardware purchased separately that had Linux installed on it later.
As of 2024, estimates suggest Linux accounts for at least 80% of the public cloud workload, partly thanks to its widespread use in platforms like Amazon Web Services (AWS), Microsoft Azure, and Google Cloud Platform.[^108][^109][^110]
ZDNet reports that 96.3% of the top one million web servers are running Linux.[^111][^110]
W3Techs states that Linux powers at least 39.2% of websites whose operating system is known, with other estimates saying 55%.[^111][^112]
Mobile devices

Android, which is based on the Linux kernel, has become the dominant operating system for smartphones. In April 2023, 68.61% of mobile devices accessing websites using [[StatCounter]] were from Android.[^113] Android is also a popular operating system for tablets, being responsible for more than 60% of tablet sales as of 2013.[^114] According to web server statistics, as of October 2021 Android has a market share of about 71%, with [[IOS]] holding 28%, and the remaining 1% attributed to various niche platforms.[^115]
Film production
For years, Linux has been the platform of choice in the film industry. The first major film produced on Linux servers was 1997's *[[Titanic (1997 film)|Titanic]]*.[^116][^117] Since then major studios including [[DreamWorks Animation]], [[Pixar]], [[Weta Digital]], and [[Industrial Light & Magic]] have migrated to Linux.[^118][^119][^120] According to the Linux Movies Group, more than 95% of the servers and desktops at large animation and visual effects companies use Linux.[^121]
Use in government
Linux distributions have also gained popularity with various local and national governments. News of the Russian military creating its own Linux distribution has also surfaced, and has come to fruition as the G.H.ost Project.[^122] The Indian state of [[Kerala]] has gone to the extent of mandating that all state high schools run Linux on their computers.[^123][^124] [[China]] uses Linux exclusively as the operating system for its [[Loongson]] processor family to achieve technology independence.[^125] In Spain, some regions have developed their own Linux distributions, which are widely used in education and official institutions, like gnuLinEx in Extremadura and Guadalinex in Andalusia. [[France]] and [[Germany]] have also taken steps toward the adoption of Linux.[^126] In April 2026 France ordered its ministries to begin to replace 2.5 million Windows desktops with Linux.[^127]  North Korea's [[Red Star OS]], developed as of 2002, is based on a version of Fedora Linux.[^128]

## Design

Many developers of [[Free and open-source software]] agree that the Linux kernel was not designed but rather [[Evolution|evolved]] through [[Natural selection]]. Torvalds considers that although the design of Unix served as a scaffolding, "Linux grew with a lot of mutations – and because the mutations were less than random, they were faster and more directed than [[Mutation|alpha-particles in DNA]]."[^129] [[Eric S. Raymond]] considers Linux's revolutionary aspects to be social, not technical: before Linux, complex software was designed carefully by small groups, but "Linux evolved in a completely different way. From nearly the beginning, it was rather casually hacked on by huge numbers of volunteers coordinating only through the Internet. Quality was maintained not by rigid standards or autocracy but by the naively simple strategy of releasing every week and getting feedback from hundreds of users within days, creating a sort of rapid Darwinian selection on the mutations introduced by developers."[^130] [[Bryan Cantrill]], an engineer of a competing OS, agrees that "Linux wasn't designed, it evolved", but considers this to be a limitation, proposing that some features, especially those related to security, cannot be evolved into, "this is not a biological system at the end of the day, it's a software system."[^131]

A Linux-based system is a modular Unix-like operating system, deriving much of its basic design from principles established in Unix during the 1970s and 1980s. Such a system uses a [[Monolithic kernel]], the Linux kernel, which handles process control, networking, access to the [[Peripheral|peripherals]], and [[File system|file systems]]. [[Device driver|Device drivers]] are either integrated directly with the kernel or added as modules that are loaded while the system is running.[^132]

The GNU [[User space and kernel space|userland]] is a key part of most systems based on the Linux kernel, with Android being a notable exception. The [[Glibc|GNU C library]], an implementation of the [[C standard library]], works as a wrapper for the system calls of the Linux kernel necessary to the kernel-userspace interface, the [[GNU toolchain|toolchain]] is a broad collection of programming tools vital to Linux development (including the [[GNU Compiler Collection|compilers]] used to build the Linux kernel itself), and the [[GNU Core Utilities|coreutils]] implement many basic [[List of Unix commands|Unix tools]]. The GNU Project also develops [[Bash (Unix shell)|Bash]], a popular [[Command-line interface|CLI]] shell. The [[Graphical user interface]] (or GUI) used by most Linux systems is built on top of an implementation of the [[X Window System]].[^133] More recently, some of the Linux community has sought to move to using [[Wayland (protocol)|Wayland]] as the display server protocol, replacing X11.[^134][^135]

Many other open-source software projects contribute to Linux systems.

- [[User space and kernel space|userland]]
- [[User space and kernel space|kernel space]]
- [[Bash (Unix shell)|bash]]
- [[LibreOffice]]
- [[GIMP]]
- [[Blender (software)|Blender]]
- [[0 A.D. (video game)|0 A.D.]]
- [[Firefox|Mozilla Firefox]]
- [[Init|init daemon]]
- [[OpenRC]]
- [[Runit]]
- [[Systemd]]
- [[Daemon (computing)|System daemons]]
- [[Polkit|polkitd]]
- [[Samba (software)|smbd]]
- [[OpenSSH|sshd]]
- [[Udevd]]
- [[Windowing system]]
- [[X Window System|X11]]
- [[Wayland (protocol)|Wayland]]
- [[SurfaceFlinger]]
- [[Mesa (computer graphics)|Mesa]]
- [[AMD Software]]
- [[GTK]]
- [[Qt (software)|Qt]]
- [[Enlightenment Foundation Libraries|EFL]]
- [[Simple DirectMedia Layer|SDL]]
- [[Simple and Fast Multimedia Library|SFML]]
- [[FLTK]]
- [[GNUstep]]
- [[C standard library]]
- [[Fopen]]
- [[Exec (system call)|execv]]
- [[Malloc]]
- [[Memcpy]]
- [[Localtime]]
- [[Pthread create|pthread_create]]
- [[Subroutine|subroutines]]
- [[Glibc]]
- [[Musl]]
- [[UClibc]]
- [[Bionic (software)|bionic]]
- [[Android (operating system)|Android]]
- [[POSIX]]
- [[Single UNIX Specification|SUS]]
- [[Linux Kernel]]
- [[Stat (system call)|stat]]
- [[Splice (system call)|splice]]
- [[Dup (system call)|dup]]
- [[Read (system call)|read]]
- [[Open (system call)|open]]
- [[Ioctl]]
- [[Write (system call)|write]]
- [[Mmap]]
- [[Close (system call)|close]]
- [[Exit (system call)|exit]]
- [[System call|System Call Interface]]
- [[Scheduling (computing)|Process scheduling]]
- [[Inter-process communication|IPC]]
- [[Memory management]]
- [[Virtual file system|Virtual files]]
- [[Advanced Linux Sound Architecture|ALSA]]
- [[Direct Rendering Infrastructure|DRI]]
- [[Evdev]]
- [[Klibc]]
- [[Logical Volume Manager (Linux)|LVM]]
- [[Device mapper]]
- [[Linux Network Scheduler]]
- [[Netfilter]]
- [[Linux Security Modules]]
- [[Security-Enhanced Linux|SELinux]]
- [[TOMOYO Linux|TOMOYO]]
- [[AppArmor]]
- [[Smack (Linux security module)|Smack]]
- [[Central processing unit|CPU]]
- [[Random-access memory|main memory]]
- [[Computer data storage|data storage devices]]

Installed components of a Linux system include the following:[^133][^137]

- A [[Bootloader]], for example [[GNU GRUB]], [[LILO (bootloader)|LILO]], [[SYSLINUX]] or [[Systemd-boot]]. This is a program that loads the Linux kernel into the computer's [[Main memory]], by being executed by the computer when it is turned on and after the [[Firmware]] initialization is performed.
- An [[Init]] program, such as the traditional [[Sysvinit]] and the newer [[Systemd]], [[OpenRC]] and [[Upstart (software)|Upstart]]. This is the first [[Process (computing)|process]] launched by the Linux kernel, and is at the root of the process tree. It starts processes such as system services and login prompts (whether graphical or in terminal mode).
- [[Library (computing)|Software libraries]], which contain code that can be used by running processes. On Linux systems using [[Executable and Linkable Format|ELF]]-format executable files, the [[Dynamic linker]] that manages the use of dynamic libraries is known as [[Ld-linux.so]]. If the system is set up for the user to compile software themselves, [[Header file|header files]] will also be included to describe the [[API|programming interface]] of installed libraries. Besides the most commonly used software library on Linux systems, the [[GNU C Library]] (glibc), there are numerous other libraries, such as [[Simple DirectMedia Layer|SDL]] and [[Mesa (computer graphics)|Mesa]].

The [[C standard library]] is the library necessary to run programs written in [[C (programming language)|C]] on a computer system, with the GNU C Library being the standard. It provides an implementation of the POSIX API, as well as extensions to that API. For embedded systems, alternatives such as [[Musl]], [[EGLIBC]] (a glibc fork once used by Debian) and [[UClibc]] (which was designed for [[UClinux]]) have been developed, although the last two are no longer maintained. Android uses its own C library, [[Bionic (software)|Bionic]]. However, musl can additionally be used as a replacement for glibc on desktop and laptop systems, as seen on certain Linux distributions like [[Void Linux]].
- Basic Unix commands, with GNU coreutils being the standard implementation. Alternatives exist for embedded systems, such as the copyleft [[BusyBox]], and the BSD-licensed [[Toybox]].
- [[Widget toolkit|Widget toolkits]] are the libraries used to build [[Graphical user interface|graphical user interfaces]] (GUIs) for software applications. Numerous widget toolkits are available, including [[GTK]] and Clutter developed by the [[GNOME Project]], [[Qt (software)|Qt]] developed by the [[Qt Project]] and led by [[The Qt Company]], and [[Enlightenment Foundation Libraries]] (EFL) developed primarily by the [[Enlightenment (software)|Enlightenment]] team.
- A [[Package manager|package management system]], such as [[Dpkg]] and [[RPM Package Manager|RPM]]. Alternatively packages can be compiled from binary or source [[Tar (computing)|tarballs]].
- User interface programs such as command shells or windowing environments.

### User interface

The [[User interface]], also known as the [[Shell|shell]], is either a command-line interface (CLI), a graphical user interface (GUI), or controls attached to the associated hardware, which is common for embedded systems. For desktop systems, the default user interface is usually graphical, although the CLI is commonly available through [[Terminal emulator]] windows or on a separate [[Virtual console]].

CLI shells are text-based user interfaces, which use text for both input and output. The dominant shell used in Linux is the [[Bash (Unix shell)|Bourne-Again Shell]] (bash), originally developed for the GNU Project; [[List of command-line interpreters|other shells]] such as [[Z shell|Zsh]] are also used.[^138][^139] Most low-level Linux components, including various parts of the [[Userland (computing)|userland]], use the CLI exclusively. The CLI is particularly suited for automation of repetitive or delayed tasks and provides very simple [[Inter-process communication]].

On desktop systems, the most popular user interfaces are the [[GUI shell|GUI shells]], packaged together with extensive [[Desktop environment|desktop environments]], such as [[KDE Plasma]], [[GNOME]], [[MATE (desktop environment)|MATE]], [[Cinnamon (desktop environment)|Cinnamon]], [[LXDE]], [[Elementary OS|Pantheon]], and [[Xfce]], though a variety of additional user interfaces exist. Most popular user interfaces are based on the X Window System, often simply called "X" or "X11". It provides [[Network transparency]] and permits a graphical application running on one system to be displayed on another where a user may interact with the application; however, certain extensions of the X Window System are not capable of working over the network.[^140] Several X display servers exist, with the reference implementation, [[X.Org Server]], being the most popular.

Several types of [[Window manager|window managers]] exist for X11, including [[Tiling window manager|tiling]], [[Dynamic window manager|dynamic]], [[Stacking window manager|stacking]], and [[Compositing window manager|compositing]]. Window managers provide means to control the placement and appearance of individual application windows, and interact with the X Window System. Simpler [[X window manager|X window managers]] such as [[Dwm]], [[Ratpoison]], or [[I3 (window manager)|i3wm]] provide a [[Minimalism (computing)|minimalist]] functionality, while more elaborate window managers such as [[FVWM]], [[Enlightenment (software)|Enlightenment]], or [[Window Maker]] provide more features such as a built-in [[Taskbar]] and [[Theme (computing)|themes]], but are still lightweight when compared to desktop environments. Desktop environments include window managers as part of their standard installations, such as [[Mutter (software)|Mutter]] (GNOME), [[KWin]] (KDE), or [[Xfwm]] (xfce), although users may choose to use a different window manager if preferred.

Wayland is a display server protocol intended as a replacement for the X11 protocol; as of 2022, it has received relatively wide adoption.[^141] Unlike X11, Wayland does not need an external window manager and compositing manager. Therefore, a Wayland compositor takes the role of the display server, window manager, and compositing manager. Weston is the reference implementation of Wayland, while GNOME's Mutter and KDE's KWin are being ported to Wayland as standalone display servers. Enlightenment has already been successfully ported since version 19.[^142] Additionally, many window managers have been made for Wayland, such as Sway or Hyprland, as well as other graphical utilities such as Waybar or Rofi.

### Video input infrastructure

Linux currently has two modern kernel-userspace APIs for handling video input devices: [[Video4Linux|V4L2]] API for video streams and radio, and [[DVB]] API for digital TV reception.[^143]

The diversity of hardware standards and data formats requires continuous architectural refinement to maintain compatibility across different device classes. Development focuses on providing robust userspace libraries to abstract device complexity, ensuring that applications can interface with varying formats without specialized hardware-specific drivers.[^144][^145]

## Development

![image|560](https://upload.wikimedia.org/wikipedia/commons/c/cd/Unix_timeline.en.svg)
*Simplified history of Unix-like operating systems. Linux shares similar architecture and concepts (as part of the [[POSIX]] standard) but does not share non-free source code with the original [[Unix]] or Minix.*

The primary difference between Linux and many other popular contemporary operating systems is that the Linux kernel and other components are free and open-source software. Linux is not the only such operating system, although it is by far the most widely used.[^146] Some [[Free software license|free]] and [[Open-source license|open-source software licenses]] are based on the principle of [[Copyleft]], a kind of reciprocity: any work derived from a copyleft piece of software must also be copyleft itself. The most common free software license, the GNU General Public License (GPL), is a form of copyleft and is used for the Linux kernel and many of the components from the GNU Project.[^147]

Linux-based distributions are intended by developers for [[Interoperability]] with other operating systems and established computing standards. Linux systems adhere to POSIX,[^148] [[Single UNIX Specification]] (SUS),[^149] [[Linux Standard Base]] (LSB), [[International Organization for Standardization|ISO]], and [[American National Standards Institute|ANSI]] standards where possible, although to date only one Linux distribution has been POSIX.1 certified, Linux-FT.[^150][^151] [[The Open Group]] has tested and certified at least two Linux distributions as qualifying for the Unix trademark, [[EulerOS]] and [[Inspur K-UX]].[^152]

Free software projects, although developed through [[Collaboration]], are often produced independently of each other. The fact that the software licenses explicitly permit redistribution, however, provides a basis for larger-scale projects that collect the software produced by stand-alone projects and make it available all at once in the form of a Linux distribution.

Many Linux distributions manage a remote collection of system software and application software packages available for download and installation through a network connection. This allows users to adapt the operating system to their specific needs. Distributions are maintained by individuals, loose-knit teams, volunteer organizations, and commercial entities. A distribution is responsible for the default configuration of the installed Linux kernel, general system security, and more generally integration of the different software packages into a coherent whole. Distributions typically use a package manager such as [[APT (software)|apt]], [[Yum (software)|yum]], [[ZYpp|zypper]], [[Pacman (package manager)|pacman]] or [[Portage (software)|portage]] to install, remove, and update all of a system's software from one central location.[^153]

### Community

![image|250](https://upload.wikimedia.org/wikipedia/commons/thumb/a/a6/Gnucash_2_screenshot.png/250px-Gnucash_2_screenshot.png?utm_source=en.wikipedia.org&utm_campaign=parser&utm_content=thumbnail)
*Photo of GNUCash accounting and bookkeeping program*

A distribution is largely driven by its developer and user communities. Some vendors develop and fund their distributions on a volunteer basis, [[Debian]] being a well-known example. Others maintain a community version of their commercial distributions, as [[Red Hat]] does with [[Fedora Linux|Fedora]], and [[SUSE S.A.|SUSE]] does with [[OpenSUSE]].[^154][^155]

In many cities and regions, local associations known as [[Linux User Group|Linux User Groups]] (LUGs) seek to promote their preferred distribution and by extension free software. They hold meetings and provide free demonstrations, training, technical support, and operating system installation to new users. Many Internet communities also provide support to Linux users and developers. Most distributions and free software / open-source projects have [[Internet Relay Chat|IRC]] chatrooms or [[Newsgroup|newsgroups]]. [[Internet forum|Online forums]] are another means of support, with notable examples being [[Stack Exchange|Unix & Linux Stack Exchange]],[^156][^157] [[LinuxQuestions.org]] and the various distribution-specific support and community forums, such as ones for [[Ubuntu]], [[Fedora Linux|Fedora]], [[Arch Linux]], [[Gentoo Linux|Gentoo]], etc. Linux distributions host [[Mailing list|mailing lists]]; commonly there will be a specific topic such as usage or development for a given list.

There are several technology websites with a Linux focus. Print magazines on Linux often bundle [[Cover disk|cover disks]] that carry software or even complete Linux distributions.[^158][^159]

Although Linux distributions are generally available without charge, several large corporations sell, support, and contribute to the development of the components of the system and free software. An analysis of the Linux kernel in 2017 showed that well over 85% of the code was developed by programmers who are being paid for their work, leaving about 8.2% to unpaid developers and 4.1% unclassified.[^160] Some of the major corporations that provide contributions include [[Intel]], [[Samsung]], [[Google]], [[AMD]], [[Oracle Corporation|Oracle]], and [[Facebook]].[^160] Several corporations, notably Red Hat, [[Canonical (company)|Canonical]], and [[SUSE S.A.|SUSE]] have built a significant business around Linux distributions.

The [[Free software license|free software licenses]], on which the various software packages of a distribution built on the Linux kernel are based, explicitly accommodate and encourage commercialization; the relationship between a Linux distribution as a whole and individual vendors may be seen as. One common Business model of commercial suppliers is charging for support, especially for business users. A number of companies also offer a specialized business version of their distribution, which adds proprietary support packages and tools to administer higher numbers of installations or to simplify administrative tasks.[^161]

Another business model is to give away the software to sell hardware. This used to be the norm in the computer industry, with operating systems such as [[CP/M]], [[Apple DOS]], and versions of the [[Classic Mac OS]] before 7.6 freely copyable (but not modifiable). As computer hardware standardized throughout the 1980s, it became more difficult for hardware manufacturers to profit from this tactic, as the OS would run on any manufacturer's computer that shared the same architecture.[^162][^163]

### Programming on Linux

Most [[Programming Language|programming languages]] support Linux either directly or through third-party community based [[Porting|ports]].[^164] The original development tools used for building both Linux applications and operating system programs are found within the [[GNU toolchain]], which includes the [[GNU Compiler Collection]] (GCC) and the [[GNU Build System]]. Among others, GCC provides compilers for [[Ada (programming language)|Ada]], [[C (programming language)|C]], [[C++]], [[Go (programming language)|Go]] and [[Fortran]]. Many programming languages have a cross-platform reference implementation that supports Linux, for example [[PHP]], [[Perl]], [[Ruby (programming language)|Ruby]], [[Python (programming language)|Python]], [[Java (programming language)|Java]], Go, [[Rust (programming language)|Rust]] and [[Haskell]]. First released in 2003, the [[LLVM]] project provides an alternative cross-platform open-source compiler for many languages. [[Proprietary software|Proprietary]] compilers for Linux include the [[Intel C++ Compiler]], [[Sun Studio (software)|Sun Studio]], and [[IBM XL C/C++ Compilers|IBM XL C/C++ Compiler]]. [[BASIC]] is available in [[Procedural programming|procedural]] form from [[QB64]], [[PureBasic]], [[Yabasic]], [[GLBasic]], [[Basic4GL]], [[XBasic]], [[WxBasic]], [[SdlBasic]], and [[Basic-256]], and in [[Object oriented programming|object oriented]] through [[Gambas]], [[FreeBASIC]], B4X, Basic for Qt, Phoenix Object Basic, [[NS Basic]], ProvideX, [[Chipmunk Basic]], [[RapidQ]] and [[Xojo]]. [[Pascal (programming language)|Pascal]] is implemented through [[GNU Pascal]], [[Free Pascal]], [[Virtual Pascal]], and graphically via [[Lazarus (software)|Lazarus]], [[PascalABC.NET]], or [[Delphi (software)|Delphi]] using [[FireMonkey]] (formerly through [[Borland Kylix]]).[^165][^166]

A common feature of Unix-like systems, Linux includes traditional specific-purpose programming languages targeted at [[Scripting language|scripting]], text processing and system configuration and management in general. Linux distributions support [[Shell script|shell scripts]], [[AWK]], [[Sed]] and [[Make (software)|make]]. Many programs also have an embedded programming language to support configuring or programming themselves. For example, [[Regular expression|regular expressions]] are supported in programs like [[grep]] and [[Locate (Unix)|locate]], the traditional Unix message transfer agent [[Sendmail]] contains its own [[Turing completeness|Turing complete]] scripting system, and the advanced text editor [[GNU Emacs]] is built around a general purpose [[Emacs Lisp|Lisp]] interpreter.[^167][^168][^169]

Most distributions also include support for [[PHP]], [[Perl]], [[Ruby (programming language)|Ruby]], [[Python (programming language)|Python]] and other [[Dynamic programming language|dynamic languages]]. While not as common, Linux also supports [[C Sharp (programming language)|C#]] and other [[Common Language Infrastructure|CLI]] [[List of CLI languages|languages]] (via [[Mono (software)|Mono]]), [[Vala (programming language)|Vala]], and [[Scheme (programming language)|Scheme]]. [[Guile (programming language)|Guile Scheme]] acts as an [[Scripting language|extension language]] targeting the GNU system utilities, seeking to make the conventionally small, [[Static typing|static]], compiled C programs of [[Unix philosophy|Unix design]] rapidly and dynamically extensible via an elegant, [[Functional programming|functional]] high-level scripting system; many GNU programs can be compiled with optional Guile [[Language binding|bindings]] to this end. Many [[Java virtual machine|Java virtual machines]] and development kits run on Linux, including the original Sun Microsystems JVM ([[HotSpot]]), IBM's J2SE RE, and many open-source projects like [[Kaffe]] and [[Jikes RVM]]; [[Kotlin]], [[Scala (programming language)|Scala]], [[Apache Groovy|Groovy]] and other [[List of JVM languages|JVM languages]] are also available.

[[GNOME]] and [[KDE]] are popular desktop environments and provide a framework for developing applications. These projects are based on the [[GTK]] and [[Qt (software)|Qt]] widget toolkits, respectively, which can also be used independently of the larger framework. Both support a wide variety of languages. There are  of [[Integrated development environment|Integrated development environments]] available including [[Anjuta]], [[Code::Blocks]], [[CodeLite]], [[Eclipse (software)|Eclipse]], [[Geany]], [[ActiveState Komodo]], [[KDevelop]], [[Lazarus (software)|Lazarus]], [[MonoDevelop]], [[NetBeans]], and [[Qt Creator]], while the long-established editors [[Vim (text editor)|Vim]], [[GNU nano|nano]] and [[Emacs]] remain popular.[^170]

### Graphics and graphics design

A topic of perennial discussion is the area of graphic design on Linux. Typical common graphics and image tools are programs such as: [[GNU Image Manipulation Program]] (GIMP), [[Inkscape]], [[Krita]], [[DaVinci Resolve]], and [[Scribus]].

## Hardware support

![image|560](https://upload.wikimedia.org/wikipedia/commons/3/3a/Linux_kernel_ubiquity.svg)
*Linux is found on many types of hardware.*

The Linux kernel is a widely ported operating system kernel, available for devices ranging from mobile phones to supercomputers; it runs on a highly diverse range of [[Computer architecture|computer architectures]], including [[ARM architecture family|ARM]]-based Android smartphones and the [[IBM Z]] mainframes. Specialized distributions and kernel forks exist for less mainstream architectures; for example, the [[Embeddable Linux Kernel Subset|ELKS]] kernel [[Fork (software development)|fork]] can run on [[Intel 8086]] or [[Intel 80286]] 16-bit microprocessors,[^171] while the [[ΜClinux]] kernel fork may run on systems without a [[Memory management unit]].[^172] The kernel also runs on architectures that were only ever intended to use a proprietary manufacturer-created operating system, such as [[Macintosh]] computers[^173][^174] (with [[PowerPC]], [[Intel]], and [[Apple silicon]] processors), [[Personal digital assistant|PDAs]], [[Video game console|video game consoles]], [[Portable media player|portable music players]], and mobile phones.

Linux has a reputation for supporting old hardware very well by maintaining standardized drivers for a long time.[^175] There are several industry associations and hardware [[Convention (meeting)|conferences]] devoted to maintaining and improving support for diverse hardware under Linux, such as [[FreedomHEC]]. Over time, support for different hardware has improved in Linux, resulting in any off-the-shelf purchase having a "good chance" of being compatible.[^176]

In 2014, a new initiative was launched to automatically collect a database of all tested hardware configurations.[^177]

## Security

One commonly expressed attribute about Linux (or open source software in general) is that through regular vetting by many individuals the software is progressively made more secure, a process referred to as being called "[[Linus's law]]".[^178][^179][^180] Linux is often highlighted for its share of the estimated global [[Usage share of operating systems|servers market share]].[^181] Besides many volunteers, the not-for-profit bodies such as the [[Open Source Security Foundation]] and [[Linux Foundation]] actively pursue ways to promote security in the Linux platform as well.

## See also

- [[Comparison of Linux distributions]]
- [[Comparison of open-source and closed-source software]]
- [[Comparison of operating systems]]
- [[Comparison of X Window System desktop environments]]
- [[Backward compatibility]]
- [[Criticism of Linux]]
- [[Linux kernel version history]]
- [[Linux Documentation Project]]
- [[Linux From Scratch]]
- [[Linux Software Map]]
- [[LinuxQuestions.org]]
- [[Linux User Group]] (LUG)
- [[List of computer books|List of Linux books]]
- [[List of free and open-source software packages]]
- [[List of live CDs]]
- [[Live USB]]
- [[List of software bugs]]
- [[List of computer security certifications]]
- [[List of computing and IT abbreviations]]
- [[List of Linux distributions]]
- [[List of Linux games|List of games released on Linux]]
- [[List of open-source mobile phones]]
- [[List of operating systems]]
- [[Loadable kernel module]]
- [[Usage share of operating systems]]
- [[Timeline of operating systems]]

- [Graphical map of Linux Internals](https://web.archive.org/web/20100211130125/http://www.makelinux.net/system/new) (archived)
- [Linux kernel website and archives](https://kernel.org)
- [The History of Linux in GIT Repository Format 1992–2010](https://archive.org/details/git-history-of-linux) (archived)


> [!note]- Notes
> 1. Cloud computing, embedded devices, mainframe computers, mobile devices, personal computers, servers, and supercomputers
> 2. Alpha, ARC, ARM, C-Sky, Hexagon, LoongArch, m68k, Microblaze, MIPS, Nios II, OpenRISC, PA-RISC, PowerPC, Power ISA, RISC-V, ESA/390, z/Architecture, SuperH, SPARC, x86, and Xtensa
> 3. util-linux is the standard set of utilities for use as part of the Linux operating system[3] in addition to one of the following userlands:
> 4. GNU Core Utilities is a userland[4][5] used by most Linux distributions.[6]
> 5. BusyBox is a userland written with size-optimization and limited resources in mind, used in many embedded Linux distributions. BusyBox replaces most GNU Core Utilities.[7]
> 6. Toybox is a userland that combines over 200 Unix command line utilities together into a single BSD-licensed executable. After a talk at the 2013 Embedded Linux Conference, Google merged toybox into AOSP and began shipping toybox in Android Marshmallow in 2015.[8]
> 7. The name "Linux" is a trademark owned by Linus Torvalds[10] and administered by the Linux Mark Institute.
> 8. As measured by the TOP500 list, which uses HPL to measure computational power
> 9. As measured by the TOP500 list, which uses HPL to measure computational power
> 10. Linux-syscall-note on GitHub
> 11. "You have to design it you cannot asymptotically reach Security." Cantrill 2017[full citation needed]


